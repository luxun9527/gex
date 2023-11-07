# go-zero 数字货币所demo

基于go-zero 开发一个数字demo,实现了现货交易的一些基本功能，现价单，市价单。

核心模块：订单，撮合，账户。

## 基本功能

### 现价单

![](http://g.recordit.co/Hh4Aa60wdd.gif)

### 市价单

![](http://g.recordit.co/JT3sxlpRQX.gif)



## 运行项目

项目依赖的中间件：消息组件pulsar，数据库mysql,redis，分布式事务dtm，websocket推送gpush。

启动项目推荐使用docker启动项目 
1、配置一个host 映射 api.gex.com 虚拟机的ip
2、启动项目
```
root@ubuntu:~/smb# docker-compose -v
Docker Compose version v2.6.1

root@ubuntu:~/smb# docker version
Client: Docker Engine - Community
 Version:           24.0.6
 API version:       1.43
 Go version:        go1.20.7
 Git commit:        ed223bc
 Built:             Mon Sep  4 12:32:12 2023
 OS/Arch:           linux/amd64
 Context:           default
 
执行命令 
make run 启动项目。
make clear 删除镜像和容器。

 
```

3、直接访问虚拟机的ip,默认nginx容器使用的是80端口





refer https://github.com/michaelliao/warpexchange/

待续。
