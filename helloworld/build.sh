#!/bin/bash
set -e

version=$1
if [ "$1" == "" ]; then
  version=latest
fi

env GOOS=linux GOARCH=amd64 go build -o app

docker build -t huyinghuan/helloworld:$version .

# if [ "$version" != "latest" ]; then
#   docker tag docker.hunantv.com/huyinghuan/zt_comment:$version docker.hunantv.com/huyinghuan/zt_comment:latest
# fi

rm app