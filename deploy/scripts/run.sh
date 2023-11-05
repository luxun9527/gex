#!/bin/bash




lang=$(cat resource/language/zh-CN.yaml)

match=$(cat app/match/rpc/etc/match-deploy.yaml)

docker-compose -f deploy/depend/docker-compose.yaml up -d

sleep 30s

docker exec -it etcd /usr/local/bin/etcdctl put language/zh-CN -- "$lang"

docker exec -it etcd /usr/local/bin/etcdctl put config/match/BTC_USDT -- "$match"

docker exec -it pulsar /pulsar/bin/pulsar-admin namespaces create public/trade

docker-compose -f deploy/dockerfiles/docker-compose.yaml up -d



