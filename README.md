# alertmanager-wechat-robot

## 项目说明
### 项目设计
- 接收alertmanager webhook，推送到企业微信群组机器人进行报警
- 通过告警标签信息，自定义配置策略，实现 @相关人员通知
- setting配置alert地址，通过API同步告警规则
- 告警历史记录、分析、dashboard进行展示
- UI页面管理，待完善

### 目录结构
```
├── protocol                       # rpc / http 功能加载
│   ├── grpc.go              
│   └── http.go    
├── cmd                            # 处理程序启停参数，加载系统配置文件
│   ├── root.go             
│   └── start.go                
├── conf                           # 配置文件加载
│   ├── config.go                  # 配置文件定义
│   ├── load.go                    # 不同的配置加载方式
│   └── log.go                     # 日志配置文件
├── etc                            # 配置文件
│   ├── config.env
│   └── config.toml
├── apps                           # 具体业务场景的领域包
│   ├── all
│   │   |-- grpc.go                # 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载。  
│   │   |-- http.go                # 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载。                    
│   ├── user                       # 具体业务场景领域服务 user
│   │   ├── api                    # http 
│   │   │    ├── user.go           # user 服务的http方法实现
│   │   │    └── http.go           # 领域模块内的 http 路由处理，向系统层注册http服务
│   │   ├── impl                   # rpc
│   │   │    ├── user.go           # user 服务的rpc方法实现
│   │   │    └── impl.go           # 领域模块内的 rpc 服务注册 ，向系统层注册rpc服务
│   │   ├──  pb                    # protobuf 定义
│   │   │     └── user.proto       # user proto 定义文件
│   │   ├── user.go                # user app 只定义扩展
│   │   ├── user.pb.go             # protobuf 生成的文件
│   │   └── user_grpc.pb.go        # pb/user.proto 生成方法定义
├── version                        # 程序版本信息
│   └── version.go
├── deploy                         # 服务部署
│   ├── docker-compose.yaml        # docker 
│   ├── deployment.yaml            # kubernetes
│   └── Dockerfile                 # 镜像制作
├── register
│   ├── grpc.go                    # rpc 服务注册
│   ├── restful.go                 # http 服务注册
│   └── all.go                     # 初始化 服务
├── README.md                    
├── main.go                        # Go程序唯一入口
├── Makefile                       # make 命令定义
└── go.mod                         # go mod 依赖定义
```

## 启动服务
### mongo创建
```
use monitoring
db.createUser({user: "monitoring", pwd: "123456", roles: [{ role: "dbOwner", db: "monitoring" }]})
```

### 配置文件
```
[app]
name = "alertmanager-wechat-robot"

[app.http]
host = "0.0.0.0"
port = "8720"

[app.grpc]
host = "0.0.0.0"
port = "16890"

[mongodb]
endpoints = ["127.0.0.1:27017"]
username = "monitoring"
password = "123456"
database = "monitoring"
```

### 运行服务
```sh
# 编译protobuf文件, 生成代码
make gen
# 初始化应用、创建系统用户、添加wechat信息
make init
# 下载项目的依赖
make dep
# 运行程序
make run
```

## alertmanager配置
### 报警配置
```json

```

