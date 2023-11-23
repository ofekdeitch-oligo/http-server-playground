package handlers

import (
	"playground/app/modules"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandler(appModule modules.AppModule) gin.HandlerFunc {
	reg := appModule.Metrics.Registry
	h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
