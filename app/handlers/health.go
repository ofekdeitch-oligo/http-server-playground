package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := GetHealthResponse{Status: "ok"}
		ctx.JSON(http.StatusOK, data)
	}
}
