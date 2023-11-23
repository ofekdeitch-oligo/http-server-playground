package app

import (
	"net/http"

	"playground/app/handlers"
	"playground/app/metrics"
	"playground/app/modules"

	"github.com/gin-gonic/gin"
)

func registerRoutes(g *gin.Engine, appModule modules.AppModule) {
	g.GET("/health", healthHandler())
	g.POST("/increment-count", handlers.IncrementCountHandler(appModule))
	g.GET("/metrics", handlers.MetricsHandler(appModule))
}

func healthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := GetHealthResponse{Status: "ok"}
		ctx.JSON(http.StatusOK, data)
	}
}

func Start() {
	g := SetupServer()

	g.Run(":8080")
}

func SetupServer() *gin.Engine {
	gin.EnableJsonDecoderUseNumber()
	g := gin.New()
	g.Use(gin.Recovery())

	metricsModule := metrics.New()

	module := modules.AppModule{Metrics: metricsModule}
	registerRoutes(g, module)

	return g
}
