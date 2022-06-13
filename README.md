# knife4gf
knife4gf is swagger ui for GoFrame V2, it based on knife4j.

> 特色：加载`nacos`配置

### 使用
1. 安装`knife4gf`
```
go get -u github.com/iasuma/knife4gf@{tag}
```

2. 引入`knife4gf`
```
package main

import (
    "github.com/gogf/gf/v2/frame/g"
    "github.com/iasuma/knife4gf"
)

func main() {
    s := g.Server()
    s.Plugin(&knife4gf.Knife4gf{})
    s.SetPort(8199)
    s.Run()
}
```
Start your server and visit: http://127.0.0.1:8199/kdoc/

### 鸣谢
1. [GoFrame](https://github.com/gogf/gf) is a modular, powerful, high-performance and enterprise-class application development framework of Golang 
2. [knife4j](https://gitee.com/xiaoym/knife4j) [（国外访问）](https://github.com/xiaoymin/swagger-bootstrap-ui) 是为Java MVC框架集成Swagger生成Api文档的工具,前身是swagger-bootstrap-ui
