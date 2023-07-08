#!/bin/bash
set -v

brctl addbr br-pool0
ifconfig br-pool0 up

cat <<EOF>clab.yaml | clab deploy -t clab.yaml -
name: cilium-l2-aware-lb 
topology:
  nodes:
    gw0:
      kind: linux
      image: 192.168.2.100:5000/vyos/vyos:1.4.7
      cmd: /sbin/init
      binds:
        - /lib/modules:/lib/modules
        - ./startup-conf/gw0-boot.cfg:/opt/vyatta/etc/config/config.boot
 
    br-pool0:
      kind: bridge

    server1:
      kind: linux
      image: 192.168.2.100:5000/nettool
      network-mode: container:cilium-l2-aware-lb-control-plane
      exec:
      - ip addr add 12.1.5.10/24 dev net0
      - ip route add 0.0.0.0/0 via 12.1.5.1 table 100
      - ip rule add from 12.1.5.0/24 table 100

    server2:
      kind: linux
      image: 192.168.2.100:5000/nettool
      network-mode: container:cilium-l2-aware-lb-worker
      exec:
      - ip addr add 12.1.5.11/24 dev net0
      - ip route add 0.0.0.0/0 via 12.1.5.1 table 100
      - ip rule add from 12.1.5.0/24 table 100

    server3:
      kind: linux
      image: 192.168.2.100:5000/nettool
      network-mode: container:cilium-l2-aware-lb-worker2
      exec:
      - ip addr add 12.1.5.12/24 dev net0
      - ip route add 0.0.0.0/0 via 12.1.5.1 table 100
      - ip rule add from 12.1.5.0/24 table 100


  links:
    - endpoints: ["br-pool0:br-pool0-net0", "server1:net0"]
    - endpoints: ["br-pool0:br-pool0-net1", "server2:net0"]
    - endpoints: ["br-pool0:br-pool0-net2", "server3:net0"]

    - endpoints: ["gw0:eth1", "br-pool0:br-pool0-net3"]

EOF
