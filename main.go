package main

import (
	"github.com/baburay23/src/handlers"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

//all my get routes are here for each activity
//The endpoint (or route) is the url you request for. It follows this structure
func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", handlers.Homepage)
	r.GET("/education", handlers.GetEducation)
	r.GET("/charity", handlers.GetCharity)
	r.GET("/cooking", handlers.GetCooking)
	r.GET("/relaxation", handlers.GetRelaxation)
	r.GET("/music", handlers.GetMusic)
	return r
}
