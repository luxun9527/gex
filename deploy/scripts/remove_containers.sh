#!/bin/bash

# 获取所有容器的ID
CONTAINER_IDS=$(docker ps -aq)

# 删除每个容器ID对应的容器
for CONTAINER_ID in $CONTAINER_IDS; do
    docker rm -f $CONTAINER_ID
done
