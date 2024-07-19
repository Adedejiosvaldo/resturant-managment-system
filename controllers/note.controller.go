package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})
}
func GetANote(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func CreateANote(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func UpdateANote(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func DeleteANote(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
