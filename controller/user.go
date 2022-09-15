package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hailong09/GoYelpCampAPI/database"
	"github.com/hailong09/GoYelpCampAPI/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.GetCollection(database.DB, "user")

func Register(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.User

	defer cancel()

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "error",
			"error": err.Error(),
		})
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
		})
		return
	}

	fmt.Println("Register user",user)

	result, err := userCollection.InsertOne(c, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "error",
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"message": "success",
		"data": result,
	})



}