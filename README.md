# Go_React
## API文档
## 后端线上项目部署目录结构：
```bash
blog-project/
├── api/ # 存放 API 相关的代码，如路由、控制器、中间件等
│   ├── controller/ # 存放控制器，负责处理请求和响应
│   ├── middleware/ # 存放中间件，负责处理请求前后的逻辑，如认证、日志、错误处理等
│   └── router/ # 存放路由，负责将请求分发给对应的控制器
├── cmd/ # 存放 main 函数，负责启动应用程序
│   └── blog/ # 存放博客应用的 main 函数
│       └── main.go
├── configs/ # 存放配置文件，如数据库、日志、API 等的配置
│   └── config.yaml
├── deployments/ # 存放部署相关的文件，如 Dockerfile、k8s 配置等
├── docs/ # 存放文档，如 API 文档、设计文档等
├── go.mod # Go Modules 的配置文件
├── go.sum # Go Modules 的依赖列表
├── internal/ # 存放内部使用的代码，不对外暴露
│   ├── model/ # 存放数据模型，如用户、文章、评论等的结构体定义和数据库操作
│   ├── service/ # 存放业务逻辑，如用户注册、文章发布、评论过滤等
│   └── util/ # 存放通用的工具函数，如加密、验证、格式化等
├── pkg/ # 存放可以对外提供的公共代码，如常量、错误码、接口等
└── test/ # 存放测试代码，如单元测试、集成测试、压力测试等
```
