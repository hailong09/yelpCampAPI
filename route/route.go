package route
 
import (
	"github.com/gin-gonic/gin"
	"github.com/hailong09/GoYelpCampAPI/controller"
	"github.com/hailong09/GoYelpCampAPI/middleware"
)

func InitializeRoute(router *gin.Engine) {
	campGroundRoute := router.Group("/api/campgrounds")
	{
		campGroundRoute.GET("/", controller.GetCampGrounds)
		campGroundRoute.POST("/", middleware.Auth() , controller.CreateNewCampGround)
		campGroundRoute.GET("/:campgroundId", controller.GetCampGroundByID)
		reviewRoute := campGroundRoute.Group("/:campgroundId/reviews").Use(middleware.Auth())
		{
			reviewRoute.POST("/", controller.CreateNewReview )
			reviewRoute.DELETE("/:reviewId", controller.DeleteReview)
		}
	}

	authRoute := router.Group("/api/auth")
	{
		authRoute.POST("/token", controller.GenerateToken)
		authRoute.POST("/register", controller.Register)

	}
}