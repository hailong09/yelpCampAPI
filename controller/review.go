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

var reviewCollection *mongo.Collection = database.GetCollection(database.DB, "reviews")

// Create a new review from a campground post
func CreateNewReview(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var review model.Review
	campgroundId := ctx.Param("campgroundId")
	userId := ctx.Request.Header.Get("userId")

	defer cancel()
	primitiveUserId, _ := primitive.ObjectIDFromHex(userId)
	primitiveCampgroundId, _ := primitive.ObjectIDFromHex(campgroundId)

	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error",
			"error":   err.Error(),
		})

		return
	}

	newReview := model.Review{
		Body:   review.Body,
		Rating: review.Rating,
		Author: primitiveUserId,
	}

	result, err := reviewCollection.InsertOne(c, newReview)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	campgroundCollection.UpdateOne(c,
		bson.M{"_id": primitiveCampgroundId},
		bson.M{"$push": bson.M{"reviews": result.InsertedID}})

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "success",
		"data":    result,
	})

}

// Delete a review from a campground post
func DeleteReview(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10* time.Second)
	reviewId := ctx.Param("reviewId")
	campgroundId := ctx.Param("campgroundId")
	defer cancel()

	reviewObjId, _ := primitive.ObjectIDFromHex(reviewId)
	campgroundObjId, _ := primitive.ObjectIDFromHex(campgroundId)

	result, err := reviewCollection.DeleteOne(c, bson.M{"_id": reviewObjId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "error",
			"error": err.Error(),
		})
		return
	}

	if result.DeletedCount < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "error",
			"data": "User with specified ID not found!",
		})
		return
	}

	campgroundCollection.UpdateOne(c, 
		bson.M{"_id": campgroundObjId}, 
		bson.M{"$pull": bson.M{"reviews": reviewObjId}})

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "success",
		"data": "User successfully deleted!",
	})


}
