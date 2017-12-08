#!/usr/bin/env bash
# Author: Pengg
# Description: Build the golang project.

basedir="`pwd`/gops-cmd"
serverbin="gops-server"
clientbin="gops-client"
serverdir=$basedir/$serverbin
clientdir=$basedir/$clientbin

glide install

cd $serverdir && go build && go install && echo "server: $GOPATH/bin/$serverbin"
cd $clientdir && go build && go install && echo "client: $GOPATH/bin/$clientbin"

