package main

import (
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerfiles "github.com/swaggo/files"

    docs "go_practice/docs"
    "go_practice/controllers"
)

// @title Simple Comment System API
// @version 1.0
// @description This is a simple comment system API
// @host localhost:8080
// @BasePath /api/v1
func main() {
    router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	
    v1 := router.Group("/api/v1")
    {
        v1.POST("/login", controllers.Login)
        v1.GET("/comments", controllers.GetComments)
        v1.POST("/comments", controllers.PostComment)
    }

	// Swagger 相關路由
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    router.Run(":8080")
}
