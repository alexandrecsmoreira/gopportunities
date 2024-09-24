package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangeOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Change Opening",
	})
}
