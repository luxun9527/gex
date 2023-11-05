#!/bin/bash

# 获取所有tag为latest的镜像ID
IMAGE_IDS=$(docker images --filter=reference="*:latest" --format "{{.ID}}")

# 删除每个镜像ID对应的镜像
for IMAGE_ID in $IMAGE_IDS; do
    docker rmi -f $IMAGE_ID
done

