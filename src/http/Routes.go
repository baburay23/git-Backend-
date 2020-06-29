package routes

import (
	"net/http"
	//"gobackend/src/app/Main"

	"github.com/gin-gonic/gin"
)

//"net/http"
//"github.com/gin-gonic/gin"
//"fmt"

type Activity struct {
	Charity *Charity `json: "charity"`
	Education *Education `json: "education"`
	Cooking string `json: "cooking"`
	Relaxation string `json: "relaxation"`
	Music string `json: "music"`
}

type Charity struct {
	Activity string `json: "activity"`
	Accessibility int `json: "accessibility"`
	Type string `json: "type"`
	Participants int `json: "participants"`
	Price int `json: "price"`
	Key int `json: "key"`
}
type Education struct {
	Activity string `json: "activity"`
	Accessibility int `json: "accessibility"`
	Type string `json: "type"`
	Participants int `json: "participants"`
	Price int `json: "price"`
	Key int `json: "key"`
} 

func GetEducation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H  {
		"Activity": "Learn Rust",
		"Accessibility": 1,
		"Type": "Education",
		"Participants": 1,
		"Price": 1,
		"Key": 1,
})
}

func GetCharity( c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	   "Charity": "Volunteer at foodbank",
	
	
 })
}
