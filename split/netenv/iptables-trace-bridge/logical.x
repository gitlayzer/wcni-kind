iptables trace log:

[这里有一点需要注意，就是我们的包是发生在SRC出来以后，然后进DST的时候，是需要经过FORWARD的。所以这里看到的是：prerouting]

1. default FORWARD DROP iptables
root@vm23040:~# dmesg -T
[Wed Sep  6 16:54:34 2023] Thit raw.prerouting>IN=br0 OUT= PHYSIN=br-eth1 MAC=7a:be:77:8c:e4:7c:6a:38:a1:63:9b:b2:08:00 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=15698 DF PROTO=ICMP TYPE=8 CODE=0 ID=51924 SEQ=1 
[Wed Sep  6 16:54:34 2023] Thit mangle.prerouting>IN=br0 OUT= PHYSIN=br-eth1 MAC=7a:be:77:8c:e4:7c:6a:38:a1:63:9b:b2:08:00 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=15698 DF PROTO=ICMP TYPE=8 CODE=0 ID=51924 SEQ=1 
[Wed Sep  6 16:54:34 2023] Thit nat.prerouting>IN=br0 OUT= PHYSIN=br-eth1 MAC=7a:be:77:8c:e4:7c:6a:38:a1:63:9b:b2:08:00 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=15698 DF PROTO=ICMP TYPE=8 CODE=0 ID=51924 SEQ=1 
[Wed Sep  6 16:54:34 2023] Thit mangle.forward>IN=br0 OUT=br0 PHYSIN=br-eth1 PHYSOUT=br-eth2 MAC=7a:be:77:8c:e4:7c:6a:38:a1:63:9b:b2:08:00 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=15698 DF PROTO=ICMP TYPE=8 CODE=0 ID=51924 SEQ=1 
[Wed Sep  6 16:54:34 2023] Thit filter.forward>IN=br0 OUT=br0 PHYSIN=br-eth1 PHYSOUT=br-eth2 MAC=7a:be:77:8c:e4:7c:6a:38:a1:63:9b:b2:08:00 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=15698 DF PROTO=ICMP TYPE=8 CODE=0 ID=51924 SEQ=1 
root@vm23040:~# 
# we can see that after the forwar, there is no packet next. so the packet should be drop at FORWARD.


2. with FORWARD ACCEPT
iptables -A FORWARD -i br0 -p icmp -j ACCEPT
root@vm0:~# dmesg -T
[Mon Sep  4 21:35:36 2023] Thit raw.prerouting>IN=br0 OUT= PHYSIN=br-eth1 MAC=2e:05:86:f8:0e:9a:72:81:ed:66:90:02:08:00 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=7126 DF PROTO=ICMP TYPE=8 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit mangle.prerouting>IN=br0 OUT= PHYSIN=br-eth1 MAC=2e:05:86:f8:0e:9a:72:81:ed:66:90:02:08:00 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=7126 DF PROTO=ICMP TYPE=8 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit nat.prerouting>IN=br0 OUT= PHYSIN=br-eth1 MAC=2e:05:86:f8:0e:9a:72:81:ed:66:90:02:08:00 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=7126 DF PROTO=ICMP TYPE=8 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit mangle.forward>IN=br0 OUT=br0 PHYSIN=br-eth1 PHYSOUT=br-eth2 MAC=2e:05:86:f8:0e:9a:72:81:ed:66:90:02:08:00 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=7126 DF PROTO=ICMP TYPE=8 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit filter.forward>IN=br0 OUT=br0 PHYSIN=br-eth1 PHYSOUT=br-eth2 MAC=2e:05:86:f8:0e:9a:72:81:ed:66:90:02:08:00 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=7126 DF PROTO=ICMP TYPE=8 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit mangle.postrouting>IN= OUT=br0 PHYSIN=br-eth1 PHYSOUT=br-eth2 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=7126 DF PROTO=ICMP TYPE=8 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit nat.postrouting>IN= OUT=br0 PHYSIN=br-eth1 PHYSOUT=br-eth2 SRC=10.1.1.10 DST=10.1.1.11 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=7126 DF PROTO=ICMP TYPE=8 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit raw.prerouting>IN=br0 OUT= PHYSIN=br-eth2 MAC=72:81:ed:66:90:02:2e:05:86:f8:0e:9a:08:00 SRC=10.1.1.11 DST=10.1.1.10 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=39966 PROTO=ICMP TYPE=0 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit mangle.prerouting>IN=br0 OUT= PHYSIN=br-eth2 MAC=72:81:ed:66:90:02:2e:05:86:f8:0e:9a:08:00 SRC=10.1.1.11 DST=10.1.1.10 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=39966 PROTO=ICMP TYPE=0 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit mangle.forward>IN=br0 OUT=br0 PHYSIN=br-eth2 PHYSOUT=br-eth1 MAC=72:81:ed:66:90:02:2e:05:86:f8:0e:9a:08:00 SRC=10.1.1.11 DST=10.1.1.10 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=39966 PROTO=ICMP TYPE=0 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit filter.forward>IN=br0 OUT=br0 PHYSIN=br-eth2 PHYSOUT=br-eth1 MAC=72:81:ed:66:90:02:2e:05:86:f8:0e:9a:08:00 SRC=10.1.1.11 DST=10.1.1.10 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=39966 PROTO=ICMP TYPE=0 CODE=0 ID=6180 SEQ=1 
[Mon Sep  4 21:35:36 2023] Thit mangle.postrouting>IN= OUT=br0 PHYSIN=br-eth2 PHYSOUT=br-eth1 SRC=10.1.1.11 DST=10.1.1.10 LEN=84 TOS=0x00 PREC=0x00 TTL=64 ID=39966 PROTO=ICMP TYPE=0 CODE=0 ID=6180 SEQ=1 
root@vm0:~# 

