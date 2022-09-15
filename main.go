package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hailong09/GoYelpCampAPI/route"
	
)

func main() {
	router := gin.Default()
		
	route.InitializeRoute(router)

	// Run route 
	router.Run()
}