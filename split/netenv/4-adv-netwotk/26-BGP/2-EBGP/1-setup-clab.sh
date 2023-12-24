#!/bin/bash
set -v

cat <<EOF>clab.yaml | clab deploy -t clab.yaml -
name: ebgp
topology:
  nodes:
    gw1:
      kind: linux
      image: 192.168.2.100:5000/vyos/vyos:1.4.7
      cmd: /sbin/init
      binds:
        - /lib/modules:/lib/modules
        - ./startup-conf/gw1.cfg:/opt/vyatta/etc/config/config.boot

    gw2:
      kind: linux
      image: 192.168.2.100:5000/vyos/vyos:1.4.7
      cmd: /sbin/init
      exec:
      - sh -c 'tcpdump -pne -i eth12 -w eth12.cap &'
      binds:
        - /lib/modules:/lib/modules
        - ./startup-conf/gw2.cfg:/opt/vyatta/etc/config/config.boot

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
    - endpoints: ["gw1:eth12", "gw2:eth12"]
    - endpoints: ["gw1:eth20", "server1:net0"]
    - endpoints: ["gw2:eth20", "server2:net0"]

EOF
