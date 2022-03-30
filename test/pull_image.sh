#!/bin/bash

docker pull busybox
docker tag busybox gcr.io/keti-container/busybox
docker push gcr.io/keti-container/busybox