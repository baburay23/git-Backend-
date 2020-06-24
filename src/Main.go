package main

import ( "net/http"
//"gobackend/src/handler"
"gobackend/handler/activity"
"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/cooking", handler.Cooking) 
		
	
	r.Run() // localhost:8080
}
func Cooking( c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"activity" : "Cooking",
	})

}
