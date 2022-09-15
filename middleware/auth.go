package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hailong09/GoYelpCampAPI/auth"
)

func Auth() gin.HandlerFunc{
	return func(ctx *gin.Context) {

		authorizationSplit := strings.Split(ctx.GetHeader("Authorization"), " ")
		if len(authorizationSplit) < 2  {
			ctx.JSON(401, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}

		tokenString := authorizationSplit[1]

		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}
		
		claim, err := auth.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Request.Header.Add("userId", claim.UserId)
		ctx.Next()
	}
}