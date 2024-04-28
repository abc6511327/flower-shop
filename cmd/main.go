package main

import (
    "log"
    "os"

    "flower-shop/internal/controllers"
    "flower-shop/pkg/db"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // 加载环境变量
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // 初始化数据库连接
    db.InitDB()

    // 创建Gin路由器实例
    router := gin.Default()

     // 静态文件路由
     router.Static("/static", "./static")

     // 基础路由
    router.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
    })

    // 实例化控制器
    flowerController := controllers.FlowerController{}

    // 配置路由
    router.GET("/flowers", flowerController.GetFlowers)

    // 从环境变量获取端口，如果未设置，默认为8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // 默认端口
    }

    // 启动服务器
    router.Run(":" + port)
}
