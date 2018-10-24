#!/bin/bash

# launch csi rsd service
echo "[UP] csi-stateful-rsdplugin ..."
kubectl create -f ./csi-stateful-rsdplugin.yaml
sleep 1
echo "[UP] csi-stateful-rbac ..."
kubectl create -f ./csi-stateful-rbac.yaml
sleep 1
echo "[UP] csi-nodedaemon-rsdplugin ..."
kubectl create -f ./csi-nodedaemon-rsdplugin.yaml
sleep 1
echo "[UP] csi-nodedaemon-rbac ..."
kubectl create -f ./csi-nodedaemon-rbac.yaml
sleep 2

# launch app nginx service
echo "[UP] app-nginx ..."
kubectl create -f ./app-nginx.yaml
