package main

import (
	"fmt"
	"loggin/internal/handlers"
	"loggin/internal/repository"
	"loggin/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	server 	:= gin.Default()
	repo 	:= repository.NewLogRepository()
	serv	:= service.NewFileService(repo)
	handl	:= handlers.NewLogHandler(serv)

	api := server.Group("/api/v1")
	{
		api.POST("/logs", handl.GetLog)
		api.POST("/logs/order", handl.GetLogByNum)
	}

	server.GET("/api/v1/healthstatus", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "i'm alive ;D",
		})
	})

	fmt.Println("Servidor rodando em http://localhost:8000")
	server.Run(":8000")
}
