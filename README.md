# 运维作业工具
- WEB界面
- 配置文件管理
- 定时任务管理

![images](https://github.com/pompeng/gops/blob/master/C995628D-A206-4265-BA8E-F4B8D6CF600B.png)

## 功能简介
- WEB界面，操作便利，避免人工操作多台服务器的困难及风险。
- 多环境支持，集中管理多个不同环境。
- 配置文件管理，通过界面修改配置，由ETCD下发，实现管理多台服务器的配置文件
- 定时任务管理，基于TCP的RPC调用。通过界面管理linux定时任务，秒级执行，基于动态反馈负载均衡机制，实现均衡负载分布式调用，用来替代Linux自带的crontab工具

## 安装
1. 克隆项目，`cd $GOPATH/src;git clone https://github.com/pompeng/gops.git`
1. 编译 `sh setup.sh`

## 安装web
1. 克隆项目，`git clone https://github.com/pompeng/gops-frontend.git`
1. 进入项目目录，`npm install && npm run build`

## 运行
1. 创建数据库，`gops/gops-doc/scheme.sql`
1. 编写配置文件
- gops_server: `gops/gops-doc/gops-server.conf`
- gops_client: `gops/gops-doc/gops-client.conf`
- nginx: `gops/gops-doc/nginx.conf`
1. `gops_server -c gops-server.conf`
1. `gops_client -c gops-client.conf`
