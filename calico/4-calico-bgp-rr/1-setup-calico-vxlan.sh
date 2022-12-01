#!/bin/bash

set -v 
exec &>./install-log-rec.txt
scp root@192.168.2.100:/root/tools/calicoctl /usr/bin/
chmod +x /usr/bin/calicoctl
kubectl apply -f ./calico.yaml

kubectl get nodes -owide

kubectl get pods -owide -A

cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    app: calico-vxlan
  name: calico-vxlan
spec:
  selector:
    matchLabels:
      app: calico-vxlan
  template:
    metadata:
      labels:
        app: calico-vxlan
    spec:
      containers:
      - image: 192.168.2.100:5000/nettool
        name: nettoolbox
        securityContext:
          privileged: true
---
apiVersion: v1
kind: Service
metadata:
  name: serversvc
spec:
  type: NodePort
  selector:
    app: calico-vxlan
  ports:
  - name: calico-vxlan
    port: 8080
    targetPort: 80
    nodePort: 32000
EOF

ip -d link show
