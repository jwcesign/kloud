package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthzHandler(ctx *gin.Context) {
	ReturnFormattedData(ctx, http.StatusOK, "kloud analyzer is healthy", nil)
}
