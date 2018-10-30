#!/bin/bash

# step1: shutdown app nginx service
echo "[DOWN] app-nginx ..."
kubectl delete -f ./app-nginx.yaml
#kubectl create -f ./app-mount.yaml

# step2: kubectl delete pvc xxxx
# step3: kubectl delete sc xxxx
