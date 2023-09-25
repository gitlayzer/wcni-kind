#!/bin/bash
set -v

cat <<EOF>clab.yaml | clab deploy -t clab.yaml -
name: ipip
topology:
  nodes:
    linux1:
      kind: linux
      image: 192.168.2.100:5000/nettool
      exec:
      - ip a a 10.1.5.1/24 dev eth1

      - ip addr add 172.12.1.10/24 dev eth2
      - ip l a ipip0 type ipip remote 172.12.1.11 local 172.12.1.10 dev eth2
      # - ip a a 10.1.5.0/32 dev ipip0 [optional]
      - ip l s ipip0 up

      - ip r a 10.1.8.0/24 via 172.12.1.11 dev ipip0 onlink

    linux2:
      kind: linux
      image: 192.168.2.100:5000/nettool
      exec:
      - ip a a 10.1.8.1/24 dev eth1

      - ip addr add 172.12.1.11/24 dev eth2
      - ip l a ipip0 type ipip remote 172.12.1.10 local 172.12.1.11 dev eth2
      # - ip a a 10.1.8.0/32 dev ipip0 [optional]
      - ip l s ipip0 up

      - ip r a 10.1.5.0/24 via 172.12.1.10 dev ipip0 onlink
       
    server1:
      kind: linux
      image: 192.168.2.100:5000/nettool
      exec:
      - ip addr add 10.1.5.10/24 dev net0
      - ip route replace default via 10.1.5.1

    server2:
      kind: linux
      image: 192.168.2.100:5000/nettool
      exec:
      - ip addr add 10.1.8.10/24 dev net0
      - ip route replace default via 10.1.8.1


  links:
    - endpoints: ["linux1:eth1", "server1:net0"]
    - endpoints: ["linux2:eth1", "server2:net0"]
    - endpoints: ["linux1:eth2", "linux2:eth2"]

EOF
