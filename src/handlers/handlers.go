package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Activity struct {
	Activity string `json: "activity"`
	Accessibility int `json: "accessibility"`
	Type string `json: "Type_"`
	Participants int `json: "participants"`
	Price int `json: "price"`
	Link string `json: "link"`
	Key string `json: "key"`


} 


type Charity struct {
	Activity string `json: "activity"`
	Accessibility int `json: "accessibility"`
	Type string `json: "Type_"`
	Participants int `json: "participants"`
	Price int `json: "price"`
	Link string `json: "link"`
	Key string `json: "key"`

}
  type Education struct {
	Activity string `json: "activity"`
	Accessibility int `json: "accessibility"`
	Type string `json: "Type_"`
	Participants int `json: "participants"`
	Price int `json: "price"`
	Link string `json: "link"`
	Key string `json: "key"`	
	
}   

type Cooking struct {
	Activity string `json: "activity"`
	Accessibility int `json: "accessibility"`
	Type string `json: "Type_"`
	Participants int `json: "participants"`
	Price int `json: "price"`
	Link string `json: "link"`
	Key string `json: "key"`

}


type Relaxation struct {
	Activity string `json: "activity"`
	Accessibility int `json: "accessibility"`
	Type string `json: "Type_"`
	Participants int `json: "participants"`
	Price int `json: "price"`
	Link string `json: "link"`
	Key string `json: "key"`

}


func Homepage ( c *gin.Context){
	c.JSON(http.StatusOK, gin.H {
		
		"Bored API": "Homepage",
		  
	})
}
func GetEducation( c *gin.Context) {
   educationActivities := Education {
		Activity: "Learn Elm",
		Accessibility: 1,
		Type: "Education",
		Participants: 1, 
		Price: 1,
		Link: "www", 
		Key: "1",}
c.JSON(http.StatusOK, educationActivities)
      
	
}
  
func GetCooking(c *gin.Context) {
	cookingActivities := Cooking{
		Activity: "Bake a pie", 
		Accessibility: 1, 
		Type: "Cooking",
		Participants: 1, 
		Price: 1,
		Link: "ww",
		Key: "1",}
c.JSON(http.StatusOK,cookingActivities)
 
}

func GetCharity( c *gin.Context) {
	charityActivities := Charity{ 
	 Activity: "Volunteer at foodbank", 
	 Accessibility: 1, 
	 Participants: 1, 
	 Price: 1,
	 Key: "1",
	 Type: "Charity"}
	c.JSON(http.StatusOK, charityActivities)
	
	
		

} 
func GetRelaxation( c *gin.Context) {
	relaxationActivities := Charity{ 
	 Activity: "Volunteer at foodbank", 
	 Accessibility: 1, 
	 Participants: 1, 
	 Price: 1,
	 Key: "1",
	 Type: "Charity"}
	c.JSON(http.StatusOK, relaxationActivities)
	
	
		

} 





 /* 	edu := []Education { 
			{ Activity: "Learn Rust", Accessibility: 1, Type: "Education",Participants: 1, Price: 1, Key: 1,},
			{ Activity: "Learn Kotlin", Accessibility: 1, Type: "Education", Participants: 1, Price: 1, Key: 1, },
			{ Activity: "Learn Elm", Accessibility: 1, Type: "Education", Participants: 1, Price: 1, Key: 1,},
			 }
			 rand.Seed(time.Now().Unix())

			 //ok dont randomise just call activity after bored api
			 Result := fmt.Print("Bored Api", education[rand.Intn(len(edu))])
			return Result
})  */