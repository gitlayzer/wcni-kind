#/bin/bash
set -v
HTTP_INGRESS=$(kubectl get ingress basic-ingress -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
curl -v --fail -s http://"$HTTP_INGRESS"/details/1 | jq


