package handlers

import (
	"calc/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func History(context *gin.Context) {
	context.JSON(http.StatusOK, services.ShowHistory())
}
