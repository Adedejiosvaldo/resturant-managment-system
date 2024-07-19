package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTables(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})
}
func GetATable(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func CreateATable(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func UpdateATable(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func DeleteATable(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
