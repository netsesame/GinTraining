package main

import (
	"main/contrillers"
	"main/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	initializers.LoadEnvVariables()
// 	initializers.ConnectToDb()
// 	initializers.SyncDatabase()
// }

//var staticFiles embed.FS

func main() {
	r := gin.Default()
	// 将 / 路径映射到 static/src/index.html 文件
	// 定义前端目录的路由
	r.Static("/static", "./static/src")
	//r.StaticFS("/static", http.FS(staticFiles))
	r.POST("/signup", contrillers.RegisterHandler)

	r.POST("/login", contrillers.LoginHandler)
	r.GET("/validate", middleware.ReqireAuth, contrillers.Validate)
	r.GET("/exchanges", middleware.ReqireAuth, contrillers.GetExchangesHandler)

	// 定义重定向路由
	r.GET("/", func(c *gin.Context) {
		http.Redirect(c.Writer, c.Request, "/static", http.StatusMovedPermanently)
	})

	//initializers.ReceiveMessage("positions", "direct", "bnPositions")
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
