#!/bin/bash


# 检查名为"gex"的网络是否存在
network_exists=$(docker network ls --format "{{.Name}}" --filter "name=gex")
# 如果网络不存在，则创建名为"gex"的网络
if [ -z "$network_exists" ]; then
    docker network create gex
    echo "网络 gex 创建成功！"
fi



docker-compose -f deploy/depend/docker-compose.yaml up -d

sleep 30s

docker exec -it pulsar /pulsar/bin/pulsar-admin namespaces create public/trade

docker-compose -f deploy/dockerfiles/docker-compose.yaml up -d



