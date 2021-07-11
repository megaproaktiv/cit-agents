package cfncustom

import (
	"os"
	"io"
	"github.com/gin-gonic/gin"
)


func init(){
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	// Logging to a file.
	f, _ := os.Create("api.log")
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
}



