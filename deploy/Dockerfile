FROM ubuntu:bionic
#FROM alpine
LABEL maintainers="Kubernetes Authors"
LABEL description="RSDCSI Driver"

#RUN apk update && apk add nvme-cli && modprobe mlx5_core && modprobe nvme_fabrics && modprobe nvme-rdma
RUN apt update && apt install nvme-cli modprobe && modprobe mlx5_core && modprobe nvme_fabrics && modprobe nvme-rdma

COPY ./_output/rsdplugin /rsdplugin
ENTRYPOINT ["/rsdplugin"]
