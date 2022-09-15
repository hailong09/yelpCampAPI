package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hailong09/GoYelpCampAPI/auth"

	"github.com/hailong09/GoYelpCampAPI/model"
	"go.mongodb.org/mongo-driver/bson"
)

type TokenRequest struct {
	Username    string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}




func GenerateToken(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var request TokenRequest
	var user model.User
	defer cancel()

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "error",
			"error": err.Error(),
		})
		return
	}

	// Check if email exists and password is correct
	err := userCollection.FindOne(c, bson.M{"username": request.Username}).Decode(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "error",
			"error": err.Error(),
		})
		return
	}
	
	
	credentialError := user.CheckPassword(request.Password)	
	if credentialError != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"message": "unauthorized",
			"error": credentialError.Error(),
		})
		return
	}

	

	tokenString, err := auth.GenerateJWT(user.Username, user.ID.Hex())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "error",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": tokenString,
	})

}