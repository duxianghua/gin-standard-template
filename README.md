# go-standard-project

Golang project standard demo

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

注: 上述第二个链接中认为 ORM 工具常见的代码组织 models 的形式不太合理，这里可以团队内部再讨论下。

## Libraries

根据个人经验和 Github stars 数量推荐一些第三方库

- [Gin][4] for HTTP service
- [sarama][5] for Kafka message producer and consumer
- [zap][6] for log

## Makefile

考虑 Makefile 统一使用部分共有的 target，方便对接 CICD 工具或者新人的接入，如

- version -> 输出当前的版本
- build -> 执行构建

## Todo

- [ ] All projects use same Docker base image
- [x] docker-compose.yml example
- [x] Makefile example
- [ ] Use Swagger
- [ ] JSON library
- [ ] ORM library
- [ ] Tests

## References

* [Golang Project Github Repo][1]
* [Awsome Go][2]


[1]: https://github.com/golang-standards/project-layout
[2]: https://github.com/avelino/awesome-go
[3]: https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md
[3]: https://github.com/gin-gonic/gin
[4]: https://github.com/Shopify/sarama
[5]: https://github.com/uber-go/zap
[6]: https://rakyll.org/style-packages/
