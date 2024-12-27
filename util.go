package main

import "github.com/gin-gonic/gin"

func handleError(ctx *gin.Context, statusCode int, err error) bool {
	if err != nil {
		ctx.JSON(statusCode, gin.H{"msg": err.Error()})
		return true
	}

	return false
}
