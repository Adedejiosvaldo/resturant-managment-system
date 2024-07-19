package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})
}

func GetAOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})
}
func CreateAOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})
}

func UpdateAOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})
}

func DeleteAOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})
}
