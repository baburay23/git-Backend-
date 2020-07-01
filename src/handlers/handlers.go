package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


/* type Activity struct {
	Charity *Charity `json: "charity"`
	Education *Education `json: "education"`
	Cooking string `json: "cooking"`

} */


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

type Cooking struct {
	Activity string `json: "activity"`
	Accessibility int `json: "accessibility"`
	Type string `json: "type"`
	Participants int `json: "participants"`
	Price int `json: "price"`
	Key int `json: "key"`

}

/* 
  func Stock(){

	education := []struct {
	Activity string `json: "activity"`
	Accessibility int `json: "accessibility"`
	Type string `json: "type"`
	Participants int `json: "participants"`
	Price int `json: "price"`
	Key int `json: "key"`
		
	}{
		{ Activity: "Learn Rust",
		Accessibility: 1,
		Type: "Education",
		Participants: 1, 
		Price: 1, 
		Key: 1,},
		{ Activity: "Learn Elm",
		Accessibility: 1,
		Type: "Education",
		Participants: 1, 
		Price: 1, 
		Key: 1,},
		{ Activity: "Learn ",
		Accessibility: 1,
		Type: "Education",
		Participants: 1, 
		Price: 1, 
	    Key: 1,},
	}
		fmt.Println(education)


} */
	
/* func GetEducation (c *gin.Context) {
	education := []struct {
		Activity string `json: "activity"`
		Accessibility int `json: "accessibility"`
		Type string `json: "type"`
		Participants int `json: "participants"`
		Price int `json: "price"`
		Key int `json: "key"`}{
			{ Activity: "Learn Rust",
			Accessibility: 1,
			Type: "Education",
			Participants: 1, 
			Price: 1, 
			Key: 1,},
			{ Activity: "Learn Elm",
			Accessibility: 1,
			Type: "Education",
			Participants: 1, 
			Price: 1, 
			Key: 1,},
			{ Activity: "Learn ",
			Accessibility: 1,
			Type: "Education",
			Participants: 1, 
			Price: 1, 
			Key: 1,},
		}
			//fmt.Println(education)
		c.JSON(http.StatusOK, gin.H {
			"Bored API" : education,
			
		 })
		} */
			


  func GetEducation( c *gin.Context) {
	e := []Education{
		Education{ Activity: "Learn Rust",
 		Accessibility: 1,
 		Type: "Education",
 		Participants: 1, 
 		Price: 1, 
		Key: 1,},
		{ Activity: "Learn Elm",
		Accessibility: 1,
		Type: "Education",
		Participants: 1, 
		Price: 1, 
		Key: 1,},
		{ Activity: "Learn Go ",
		Accessibility: 1,
		Type: "Education",
		Participants: 1, 
		Price: 1, 
	    Key: 1,},}
     c.JSON(http.StatusOK, gin.H {
 	   "Bored API" : e,
	
 })
 }  

 func GetCooking(c *gin.Context){
	o := Cooking{
		Activity: "Bake a pie", 
		Accessibility: 1, 
		Type: "Cooking",
		Participants: 1, 
		Price: 1,
		Key: 1,}
	c.JSON(http.StatusOK, gin.H {
			"Bored API" : o.Activity,	
	})
}

func GetCharity( c *gin.Context) {
	h := Charity{ 
	Activity: "Volunteer at foodbank", 
	 Accessibility: 1, 
	 Type: "Charity",
	 Participants: 1, 
	 Price: 1,
	  Key: 1,}
	c.JSON(http.StatusOK, gin.H {
	 "Bored API" : h.Activity,
		
 })
} 

/*   func GetEducation( c *gin.Context) {
	   e := []Education {
		Education{Activity: "Learn Rust",
		Accessibility: 1,
		Type: "Education",
		Participants: 1, 
		Price: 1, 
		Key: 1,},
		{Activity: "Learn Elm", 
		Accessibility: 1, 
		Type: "Education",
		Participants: 1, 
		Price: 1, 
		Key: 1,},
	   
		//E := education
	c.JSON( http.StatusOK, gin.H {
	  "Bored API" : e.Activity,
	 }),
	
 }  
} */




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