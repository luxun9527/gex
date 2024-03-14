# go微服务实践-基于go-zero实现一个交易所。

基于go-zero 实现一个交易所现货交易的基本功能。

- 限价单，市价单的撮合。
- 基本行情(盘口，k线，tick)，以及个人订单变化的实时推送。





## 基本架构

![](https://cdn.learnku.com/uploads/images/202402/15/51993/bBrX3MgAl6.png!large)

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



## go 实践

### 基本架构设计



### dtm saga处理分布式事务。

### etcd 实现管理后台动态配置参数。

### ast语法树 自制工具生成错误码，完善gorm gen的软删除

### 日志上报到im工具企业微信，lark,tg等工具

### Docke部署