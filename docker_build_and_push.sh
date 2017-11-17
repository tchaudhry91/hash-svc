#!/bin/bash

TAG=$1
DATE=`date '+%Y%m%d%H%M%S'`

# BUILDER
docker build -f Dockerfile-builder -t hasher:temp-$TAG .
docker run -d --name=$DATE hasher:temp-$TAG
docker cp $DATE:/main ./hash-svc
docker rm -f $DATE

# FINAL IMAGE

docker build -f Dockerfile -t tchaudhry/hash-svc:$TAG .
rm hash-svc
docker push tchaudhry/hash-svc:$TAG
