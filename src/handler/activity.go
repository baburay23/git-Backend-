 package handler

import ( "net/http"
"github.com/gin-gonic/gin"
)
func Cooking( c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"activity" : "Cooking",
	})

} 