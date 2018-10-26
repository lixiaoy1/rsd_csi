FROM ubuntu:bionic
#FROM alpine
LABEL maintainers="Kubernetes Authors"
LABEL description="RSDCSI Driver"

#RUN apk update && apk add nvme-cli && modprobe mlx5_core && modprobe nvme_fabrics && modprobe nvme-rdma
RUN apt update && apt install nvme-cli

COPY ./_output/rsdplugin /rsdplugin
ENTRYPOINT ["/rsdplugin"]
