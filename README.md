# advproject
一个基于Gin框架，以DevOps开发模式搭建的项目模板。

## 项目组成

- 接口文档 [swag](https://github.com/swaggo/swag)
- 日志 [zap](https://github.com/uber-go/zap)
- 配置 [godotenv](https://github.com/joho/godotenv)
- docker容器化文件 [Dockerfile](Dockerfile)
- Jenkins Pipline [Jenkinsfile](Jenkinsfile)
- Kubernetes部署文件 [deployment](deployment-template.yaml)

**代码结构**
```
├── Dockerfile                  // Docker镜像打包文件
├── Jenkinsfile                 // Jenkins CI 定义
├── README.md
├── api
│   └── status.go               // 运行状态接口
├── config
│   ├── base.go                 // 基础配置
│   └── init.go                 // 配置模块初始化
├── deployment-template.yaml    // Kubernetes部署文件
├── docs                        // swag 接口文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
├── serializer
│   └── common.go               // 接口格式定义
├── server
│   └── router.go               // 路由
├── test
│   └── status_test.go          // 测试代码
└── util
    └── logger.go               // 日志模块
```

## 配置说明
### 基础组件
#### 接口文档
安装swag cli
```
$ go get -u github.com/swaggo/swag/cmd/swag
```

生成API文档，每次接口注释有更新都需要更新
```
$ swag init
```

接口文档地址: http://localhost:3000/swagger/index.html

#### 系统配置
默认配置硬编码在config模块中，可通过环境变量覆盖，开发时在项目目录下新建`.env`文件来快速修改环境变量。

变量名 | 说明 | 默认值 | 可选值
---|---|---|---
RUN_MODE | 测试/生产环境 | develop | develop, prod

### 监控及日志记录
#### 日志输出

### 基础设施即代码