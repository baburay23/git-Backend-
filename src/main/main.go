package main

import (
	. "github.com/baburay23/git-Backend-/src/routes"
	"github.com/gin-gonic/gin"
)



//all my get routes are here for each activity 
//The endpoint (or route) is the url you request for. It follows this structure
func main(){
	
	//fmt.Println(Education{"Learn Elm", 1, "Education", 1, 1, 1 })
	 //rand.Seed(time.Now().Unix())
	//education := Education{ Activity: "Learn Rust", Accessibility: 1, Type: "Education",Participants: 1, Price: 1, Key: 1,}
	//Education.Activity = "Learn Rust"
	 //fmt.Println("Bored API", Education.Activity)
		
    r := gin.Default()
	r.GET("/charity",GetCharity)
    //r.GET("/cooking", GetCooking)
	r.GET("/education",GetEducation) 
		
	//The HandlerFunc type is an cadapter to allow the use of ordinary functions as HTTP handlers.
	r.Run()
  }


	




//A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
//A Request represents an HTTP request received by a server or to be sent by a client.


//Gin puts request in *gin.Context type, so you can access request headers, files and etc from a *gin.Context variable.
// You should pass this type as the argument to the second argument of http methods.
//func (c *Context) Handler() HandlerFunc
/* func GetEducation( c *gin.Context) *Education {
	var mjson Education
	c.JSON(200, mjson)
	return &Education{}
	} 





	
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