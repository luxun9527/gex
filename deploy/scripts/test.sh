#!/bin/bash
lang="$(cat resource/language/zh-CN.yaml)"

docker exec -it etcd  /usr/local/bin/etcdctl put language/zh-CN -- "$lang"