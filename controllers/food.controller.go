package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFoods(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})
}
func GetAFood(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func CreateAFood(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func UpdateAMeal(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func DeleteAMeal(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
