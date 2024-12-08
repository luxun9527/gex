#!/bin/bash


# 检查名为"gex"的网络是否存在
network_exists=$(docker network ls --format "{{.Name}}" --filter "name=gex")
# 如果网络不存在，则创建名为"gex"的网络
if [ -z "$network_exists" ]; then
    docker network create gex
    echo "网络 gex 创建成功！"
fi

lang='50006: 超过最小精度
100001: 内部错误
100002: 内部错误
100003: 内部错误
100004: 参数错误
100005: 记录未找到
100006: 重复数据
100007: 内部错误
100009: 内部错误
100010: 内部错误
100011: 内部错误
100012: 验证码错误
200001: 用户不存在
200002: 用户余额不足
200003: token验证失败
200004: token到期
200005: 账户密码验证失败1
500001: 订单未找到
500002: 订单已经成交获取已经取消
500003: 市价单不允许手动取消
500004: 订单簿没有买单
500005: 订单簿没有卖单
500006: 超过币种最小精度'

coin1='coinid: 10001
coinname: IKUN
prec: 3'

coin2='coinid: 10002
coinname: USDT
prec: 5'
symbol='symbolname: IKUN_USDT
symbolid: 1
basecoinname: IKUN
basecoinid: 10001
quotecoinname: USDT
quotecoinid: 10002
baseCoinPrec: 3
quoteCoinPrec: 5'

docker-compose -f deploy/depend/docker-compose.yaml up -d

sleep 30s

docker exec -it etcd /usr/local/bin/etcdctl put language/zh-CN -- "$lang"
docker exec -it etcd /usr/local/bin/etcdctl put Coin/IKUN -- "$coin1"
docker exec -it etcd /usr/local/bin/etcdctl put Coin/USDT -- "$coin2"
docker exec -it etcd /usr/local/bin/etcdctl put Symbol/IKUN_USDT -- "$symbol"


docker exec -it pulsar /pulsar/bin/pulsar-admin namespaces create public/trade
docker exec -it pulsar /pulsar/bin/pulsar-admin topics create persistent://public/trade/match_source_IKUN_USDT
docker exec -it pulsar /pulsar/bin/pulsar-admin topics create persistent://public/trade/match_result_IKUN_USDT

docker-compose -f deploy/dockerfiles/docker-compose.yaml up -d



