#!/bin/bash
apk update && apk add git go
export GOROOT=/usr/lib/go
export GOPATH=/go
export GOBIN=/go/bin
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
export GO15VENDOREXPERIMENT=1

cd /go/src/github.com/Dataman-Cloud/zookeeper-helper/src && \
    go build && chmod 777 src && \
    cp src /zookeeper-helper

rm -rf /tmp/* /var/tmp/*
rm -f /etc/ssh/ssh_host_*
rm -rf /go
rm -rf /usr
