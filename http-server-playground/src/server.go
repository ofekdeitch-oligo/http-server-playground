package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := SetupServer()

	g.Run(":8080")
}

func SetupServer() *gin.Engine {
	gin.EnableJsonDecoderUseNumber()
	g := gin.New()

	// Used middlewares
	g.Use(gin.Recovery())

	g.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return g
}
