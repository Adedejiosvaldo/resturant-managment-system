package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInvoices(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func GetInvoice(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func CreateInvoice(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func UpdateInvoice(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
func DeleteInvoice(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hi"})

}
