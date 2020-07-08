package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActivityTypes struct {
	Activity      string `json: "activity"`
	Accessibility int    `json: "accessibility"`
	Type          string `json: "Type_"`
	Participants  int    `json: "participants"`
	Price         int    `json: "price"`
	Link          string `json: "link"`
	Key           string `json: "key"`
}

func Homepage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Bored API": "Homepage",
	})
}
func GetEducation(c *gin.Context) {
	educationActivities := ActivityTypes{
		Activity:      "Learn Elm",
		Accessibility: 1,
		Type:          "Education",
		Participants:  1,
		Price:         1,
		Link:          "www",
		Key:           "1"}
	c.JSON(http.StatusOK, educationActivities)
}

func GetCooking(c *gin.Context) {
	cookingActivities := ActivityTypes{
		Activity:      "Bake a pie",
		Accessibility: 1,
		Type:          "Cooking",
		Participants:  1,
		Price:         1,
		Link:          "ww",
		Key:           "1"}
	c.JSON(http.StatusOK, cookingActivities)

}

func GetCharity(c *gin.Context) {
	charityActivities := ActivityTypes{
		Activity:      "Volunteer at foodbank",
		Accessibility: 1,
		Participants:  1,
		Price:         1,
		Key:           "1",
		Type:          "Charity"}
	c.JSON(http.StatusOK, charityActivities)

}
func GetRelaxation(c *gin.Context) {
	relaxationActivities := ActivityTypes{
		Activity:      "Have a facial",
		Accessibility: 1,
		Participants:  1,
		Price:         1,
		Key:           "1",
		Type:          "Charity"}
	c.JSON(http.StatusOK, relaxationActivities)

}

func GetMusic(c *gin.Context) {
	musicActivities := ActivityTypes{
		Activity:      "Leant to play the trumpet",
		Accessibility: 1,
		Participants:  1,
		Price:         1,
		Key:           "1",
		Type:          "Charity"}
	c.JSON(http.StatusOK, musicActivities)

}
