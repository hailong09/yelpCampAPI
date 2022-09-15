package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hailong09/GoYelpCampAPI/database"
	"github.com/hailong09/GoYelpCampAPI/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var campgroundCollection *mongo.Collection = database.GetCollection(database.DB, "campgrounds") 

// Create a new campground
func CreateNewCampGround(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var campground model.Campground
	userId := ctx.Request.Header.Get("userId")

	defer cancel()
	
	primitiveUserId, _ := primitive.ObjectIDFromHex(userId)
	// Call BindJSON to bind the received JSON to
	// newCampground
	if err := ctx.BindJSON(&campground); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "error",
			"error": err.Error(),
		})

		return
	}



	newCampground := model.Campground{
		Title: campground.Title,
		Geometry: campground.Geometry,
		Price: campground.Price,
		Description: campground.Description,
		Location: campground.Location,
		Author: primitiveUserId,
		
	}

	result, err := campgroundCollection.InsertOne(c, newCampground)

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

// Get all available campgrounds in from the database
func GetCampGrounds(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var campgrounds []model.Campground

	defer cancel()

	results , err := campgroundCollection.Find(c, bson.M{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "error",
			"error": err.Error(),
		})
		return
	}

	defer results.Close(c)

	for results.Next(c) {
		var singleCampground model.Campground
		if err = results.Decode(&singleCampground); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"message": "error",
				"error": err.Error(),
			})
			return
		}
		campgrounds = append(campgrounds, singleCampground)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": campgrounds,
	})

}



// Get the campground based on ID
func GetCampGroundByID(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var foundCampground model.Campground
	campgroundId := ctx.Param("campgroundId")
	defer cancel()

	primitiveUserId, _ := primitive.ObjectIDFromHex(campgroundId)

	err := campgroundCollection.FindOne(c, bson.M{"_id": primitiveUserId}).Decode(&foundCampground)
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
		"data": foundCampground,
	})

	
}

