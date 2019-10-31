package main

import (
	// "encoding/json"
	// "fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"longchain.com/memoriae/profit/config"
	"longchain.com/memoriae/profit/web/controllers"
)

func main() {
	// 监听的地址
	addr := config.Addr

	app := iris.New()
	// 中间件
	app.Use(recover.New())
	app.Use(logger.New())
	// 添加路由
	controllers.New(app)
	// 运行
	app.Run(iris.Addr(addr))
}