package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Get Users"})

}
func GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Get User"})

}
func SignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Signup"})

}
func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Lgin"})

}
