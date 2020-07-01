package main

import (
	"github.com/baburay23/src/handlers"
	"github.com/gin-gonic/gin"
)

//all my get routes are here for each activity
//The endpoint (or route) is the url you request for. It follows this structure
func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	r := gin.Default()
	r.GET("/education", handlers.GetEducation)
	r.GET("/charity", handlers.GetCharity)
	return r
}

//fmt.Println(Education{"Learn Elm", 1, "Education", 1, 1, 1 })
//rand.Seed(time.Now().Unix())
//education := Education{ Activity: "Learn Rust", Accessibility: 1, Type: "Education",Participants: 1, Price: 1, Key: 1,}
//Education.Activity = "Learn Rust"
//fmt.Println("Bored API", Education.Activity)

/* func Charity( c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"type" : "Charity",
	})
}

func Relaxation( c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"type" : "Relaxation",
	})
}

func Music( c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"type" : "Music",
	})
}  */
//}
