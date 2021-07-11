package cfncustom

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// ApiRouter
// Serves the to be mocked API
func ApiRouter() http.Handler {

	e := gin.New()
	e.Use(gin.Recovery())
	e.PUT("/", func(c *gin.Context) {
		
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"status": "sucess",
				"msg": "API",

			},
		)
	})

	return e
}