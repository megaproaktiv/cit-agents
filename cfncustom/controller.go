package cfncustom

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ControllerRouter() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"msg": "Controller API",
			},
		)
	})

	return e
}
