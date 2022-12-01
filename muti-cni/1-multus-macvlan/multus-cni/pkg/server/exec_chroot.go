// Copyright (c) 2022 Multus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/containernetworking/cni/pkg/invoke"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/version"
)

// ChrootExec implements invoke.Exec to execute CNI with chroot
type ChrootExec struct {
	Stderr     io.Writer
	chrootDir  string
	workingDir string   // working directory in the outer root
	outerRoot  *os.File // outer root directory
	version.PluginDecoder
	mu sync.Mutex
}

var _ invoke.Exec = &ChrootExec{}

func (e *ChrootExec) chroot() error {
	var err error
	e.workingDir, err = os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "getwd before chroot failed: %v\n", err)
		return fmt.Errorf("getwd before chroot failed: %v", err)
	}

	e.outerRoot, err = os.Open("/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "getwd before chroot failed: %v\n", err)
		return fmt.Errorf("getwd before chroot failed: %v", err)
	}

	if err := syscall.Chroot(e.chrootDir); err != nil {
		fmt.Fprintf(os.Stderr, "chroot to %s failed: %v\n", e.chrootDir, err)
		return fmt.Errorf("chroot to %s failed: %v", e.chrootDir, err)
	}

	if err := os.Chdir("/"); err != nil {
		fmt.Fprintf(os.Stderr, "chdir to \"/\" failed: %v\n", err)
		return fmt.Errorf("chdir to \"/\" failed: %v", err)
	}

	return nil
}

func (e *ChrootExec) escape() error {
	if e.outerRoot == nil || e.workingDir == "" {
		return nil
	}

	// change directory to outer root and close it
	if err := syscall.Fchdir(int(e.outerRoot.Fd())); err != nil {
		fmt.Fprintf(os.Stderr, "changing directory to outer root failed: %v\n", err)
		return fmt.Errorf("changing directory to outer root failed: %v", err)
	}

	if err := e.outerRoot.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "closing outer root failed: %v\n", err)
		return fmt.Errorf("closing outer root failed: %v", err)
	}

	// chroot to current directory aka "." being the outer root
	if err := syscall.Chroot("."); err != nil {
		fmt.Fprintf(os.Stderr, "chroot to current directory failed: %v\n", err)
		return fmt.Errorf("chroot to current directory failed: %v", err)
	}

	if err := os.Chdir(e.workingDir); err != nil {
		fmt.Fprintf(os.Stderr, "chdir to working directory failed: %v\n", err)
		return fmt.Errorf("chdir to working directory failed: %v", err)
	}
	e.outerRoot = nil
	e.workingDir = ""

	return nil
}

// ExecPlugin executes CNI plugin with given environment/stdin data.
func (e *ChrootExec) ExecPlugin(ctx context.Context, pluginPath string, stdinData []byte, environ []string) ([]byte, error) {
	// lock and do chroot to execute plugin with host root
	e.mu.Lock()
	defer e.mu.Unlock()
	err := e.chroot()
	defer e.escape()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ExecPlugin failed at chroot: %v\n", err)
		return nil, fmt.Errorf("ExecPlugin failed at chroot: %v", err)
	}

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	c := exec.CommandContext(ctx, pluginPath)
	c.Env = environ
	c.Stdin = bytes.NewBuffer(stdinData)
	c.Stdout = stdout
	c.Stderr = stderr

	// Retry the command on "text file busy" errors
	for i := 0; i <= 5; i++ {
		err = c.Run()

		// Command succeeded
		if err == nil {
			break
		}

		// If the plugin is currently about to be written, then we wait a
		// second and try it again
		if strings.Contains(err.Error(), "text file busy") {
			time.Sleep(time.Second)
			continue
		}

		// All other errors except than the busy text file
		return nil, e.pluginErr(err, stdout.Bytes(), stderr.Bytes())
	}

	// Copy stderr to caller's buffer in case plugin printed to both
	// stdout and stderr for some reason. Ignore failures as stderr is
	// only informational.
	if e.Stderr != nil && stderr.Len() > 0 {
		_, _ = stderr.WriteTo(e.Stderr)
	}
	return stdout.Bytes(), nil
}

func (e *ChrootExec) pluginErr(err error, stdout, stderr []byte) error {
	emsg := types.Error{}
	if len(stdout) == 0 {
		if len(stderr) == 0 {
			emsg.Msg = fmt.Sprintf("netplugin failed with no error message: %v", err)
		} else {
			emsg.Msg = fmt.Sprintf("netplugin failed: %q", string(stderr))
		}
	} else if perr := json.Unmarshal(stdout, &emsg); perr != nil {
		emsg.Msg = fmt.Sprintf("netplugin failed but error parsing its diagnostic message %q: %v", string(stdout), perr)
	}
	return &emsg
}

// FindInPath try to find CNI plugin based on given path
func (e *ChrootExec) FindInPath(plugin string, paths []string) (string, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	err := e.chroot()
	defer e.escape()
	if err != nil {
		fmt.Fprintf(os.Stderr, "FindInPath failed at chroot: %v\n", err)
		return "", fmt.Errorf("FindInPath failed at chroot: %v", err)
	}

	return invoke.FindInPath(plugin, paths)
}
