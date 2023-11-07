#!/bin/bash


# 检查名为"gex"的网络是否存在
network_exists=$(docker network ls --format "{{.Name}}" --filter "name=gex")
# 如果网络不存在，则创建名为"gex"的网络
if [ -z "$network_exists" ]; then
    docker network create gex
    echo "网络 gex 创建成功！"
fi

lang=$(cat resource/language/zh-CN.yaml)

match=$(cat app/match/rpc/etc/match-deploy.yaml)

docker-compose -f deploy/depend/docker-compose.yaml up -d

sleep 30s

docker exec -it etcd /usr/local/bin/etcdctl put language/zh-CN -- "$lang"

docker exec -it etcd /usr/local/bin/etcdctl put config/match/BTC_USDT -- "$match"

docker exec -it pulsar /pulsar/bin/pulsar-admin namespaces create public/trade

docker-compose -f deploy/dockerfiles/docker-compose.yaml up -d



