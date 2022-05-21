## Community Backend

使用`Go`语言编写的跨平台高性能简单社区APP后端，启动后访问 http://localhost:1016

完整接口文档 -> [docs.md](docs.md)

Swagger文档 -> [Swagger.json](Swagger.json)

后端处理请求时间保证在 `80ms` 一下(预热，不包括冷启动时间)

技术栈：[Sqlite](https://www.sqlite.org/index.html), [Gin](https://github.com/gin-gonic/gin), [Gorm](https://gorm.io/index.html)
