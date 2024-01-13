# 基于go-zero实现一个数字货币交易所。



# v1.1.0

基于go-zero 开发一个数字货币交易所demo,实现了交易所现货交易的一些基本功能。

- 限价单，市价单的撮合。
- 基本行情(盘口，k线，tick)，以及个人订单变化的实时推送。

核心模块：订单，撮合，账户，行情。

## 基本功能

### 限价单
![](https://s1.locimg.com/2023/11/08/10dcdafd0ae03.gif)



### 市价单

![](https://s1.locimg.com/2023/11/08/5f83f2de9742e.gif)



## 运行项目

项目依赖的中间件：消息组件pulsar，数据库mysql,redis，分布式事务dtm，websocket推送gpush。



1、配置一个host 映射， api.gex.com:项目地址

2、项目已经整理好docker-compose文件，依赖和程序分别在不同的docker-compose文件，使用docker-compose即可一键启动项目，docker版本不能太旧具体如下。

```shell
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
make build 编译项目。
make run 启动项目。
make clear 删除镜像和容器（会删除所有的容器和新建的镜像。）
账号zhangsan 密码test
账号lisi     密码test 
 
```

3、直接访问启动项目机器的ip， 默认nginx配置的是的是80端口。





初步完成了一个demo,实现了基本现货交易的一些基本功能，还有很多地方不完善和考虑不周全的地方。

项目地址 https://github.com/luxun9527/gex  如果觉得对您有帮助，您的一个star就是我更新的动力。

参考 https://github.com/michaelliao/warpexchange/



# v1.2.0

待实现

1、重构前端。

2、完成管理后台。

​	2.1 交易对管理，动态配置。

​	2.2  错误码管理。

3、 使用交易所的SDK获取价格模拟下单

4、各个服务适配动态配置

5、加入jaeger链路追踪

6、订单，用户，分库分表

7、项目更详细的介绍 ，架构设计等。

