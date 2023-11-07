#!/bin/bash
# 获取匹配的镜像列表
image_list=$(docker image ls --format "{{.Repository}}:{{.Tag}}" | grep -E "dockerfiles|depend")
# 遍历并删除每个镜像
for image in $image_list; do
    docker image rm -f  $image
done