package handlers

import (
	"github.com/gin-gonic/gin"
	"loggin/internal/dto"
	"loggin/internal/service"
	"net/http"
)

type LogHandler struct {
	service *service.LogService
}

func NewLogHandler(service *service.LogService) *LogHandler {
	return &LogHandler{
		service: service,
	}
}

func (l *LogHandler) GetLog(ctx *gin.Context) {
	var req dto.GetLogDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid payload JSON",
		})
		return
	}

	file, err := l.service.GetLogContent(req.Path)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, file)
}

func (l *LogHandler) GetLogByNum(ctx *gin.Context) {
	var req dto.GetLogByNumDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid payload JSON",
		})
		return
	}

	file, err := l.service.GetLogContentByNum(req.Path, req.Order)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, file)
}