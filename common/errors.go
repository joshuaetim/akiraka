package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(ctx *gin.Context, err error, message ...string) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": message})
		ctx.Abort()
	}
}
