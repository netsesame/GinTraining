package main

import (
	"main/contrillers"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	initializers.LoadEnvVariables()
// 	initializers.ConnectToDb()
// 	initializers.SyncDatabase()
// }

func main() {
	r := gin.Default()
	r.POST("/signup", contrillers.RegisterHandler)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
