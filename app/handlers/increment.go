package handlers

import (
	"net/http"
	"playground/app/modules"

	"github.com/gin-gonic/gin"
)

func IncrementCountHandler(appModule modules.AppModule) gin.HandlerFunc {
	countMetric := appModule.Metrics.Metrics.Count

	return func(ctx *gin.Context) {
		countMetric.Inc()
		ctx.JSON(http.StatusCreated, nil)
	}
}
