# go-standard-project

Golang project standard

## Basic

- 单个项目负责单个服务
- Semantic Versioning

## Layout

```
├── bin
├── cmd
├── docker-compose.yml
├── Dockerfile
├── docs
├── go.mod
├── go.sum
├── internal
├── Makefile
├── pkg
├── README.md
└── Version
```

- bin 存放 build 好的二进制文件
- cmd 目录用于存放项目的 entrypoint，如 main.go，如果是非服务类型 repo 也可以忽略该目录
- docs 目录用于存放项目的详细文档，简单的放在 README.md 即可
- internal 用于项目业务逻辑，代码放在该 package 外部则不可以用
- pkg 目录用于存放别的 Golang 项目可能引用的包，独立的微服务相对较少用到，一般用于提供公共服务的 repo
- docker-compose.yml 则用于服务依赖如 Redis 或者 MySQL 这类的第三方服务时编写一个完整的可启动的 docker-compose 配置
- Makefile 定义好统一的 target，方便后续配置 Pipeline

package 的分层基本原则是上层可以调用下层，每个 package 单独负责某一项功能。

更详细的说明可见 [project-layout][1] 或者 [Go 面向包的设计和架构分层][3]。