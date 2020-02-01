# advproject
一个基于Gin框架，以DevOps开发模式搭建的项目模板。

## 项目组成
> DevOps 集文化理念、实践和工具于一身，可以提高组织高速交付应用程序和服务的能力，与使用传统软件开发和基础设施管理流程相比，能够帮助组织更快地发展和改进产品。这种速度使组织能够更好地服务其客户，并在市场上更高效地参与竞争。

- 持续集成：一种软件开发实践经验，开发人员会定期将他们的代码变更合并到一个中央存储库中，之后系统会自动运行构建和测试操作
    - GitHub
    - webhook 推送 Jenkins
- 持续交付付：自动构建和测试代码更改，并为将其发布到生产环境做好准备，实现对持续集成的扩展
    - Jenkins Pipeline 测试，生成镜像，部署
- 基础设施即代码：使用基于代码的工具来连接基础设施，并且能够以处理应用程序代码的方式来处理基础设施。基础设施和服务器由代码进行定义。
    - [Dockerfile](Dockerfile) 打包Docker镜像
    - [Jenkinsfile](Jenkinsfile) 定义Jenkins Pipeline流程
    - [deployment](deployment-template.yaml) k8s Deployment 部署声明
- 监控和日志记录：组织对各项指标和日志进行监控，以了解应用程序和基础设施性能如何影响其产品的最终用户体验
    - [zap](https://github.com/uber-go/zap) 结构化日志输出，可配合k8s ELK插件使用
    - k8s livenessProbe 容器健康检查


## 配置说明
### 接口文档
安装swag cli
```
$ go get -u github.com/swaggo/swag/cmd/swag
```

生成API文档，每次接口说明有更新都需要更新
```
$ swag init
```

接口文档地址: http://localhost:3000/swagger/index.html

### 系统配置
默认配置硬编码在config模块中，可通过环境变量覆盖，开发时在项目目录下新建`.env`文件来快速修改环境变量。

变量名 | 说明 | 默认值 | 可选值
---|---|---|---
RUN_MODE | 测试/生产环境 | develop | develop, prod

### 日志输出
使用Uber开源的高性能日志框架[zap](https://github.com/uber-go/zap)来作为日志输出，直接使用了zap默认的格式配置，在develop开发环境下以易读的行模式输出，prod生产下以结构化模式输出。

### 状态接口
- 状态输出，用于定时请求进行健康检查 `/api/v1/status`
- hostname输出，在多节点部署时测试是哪个节点 `/api/v1/hostname`

### Jenkinsfile
Jenkinsfile 中镜像存储使用了腾讯云容器服务的私有镜像仓库和 [Server酱](http://sc.ftqq.com/3.version) 来推送通知，可根据实际修改。同时通过变量名来使用了一些私有密钥，需要提前在Jenkins的凭证管理中提前配置：

- tcloud-docker-reg 镜像仓库访问密钥
- kubctl-config kubectl密钥文件
- PUSH_KEY Server酱通知推送的密钥

### k8s deployment
[deployment-template.yaml](deployment-template.yaml)中定义Kubernetes部署模板，通过image指定了拉取镜像的地址，可根据需要修改。`imagePullSecrets`中使用了拉取私有镜像的config name，需要提前在k8s中配置。
