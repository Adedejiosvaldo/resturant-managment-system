package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})
}
func GetAMenu(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func CreateAMenu(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func UpdateAMenu(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func DeleteAMenu(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
