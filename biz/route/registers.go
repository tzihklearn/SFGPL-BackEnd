package route

import "github.com/gin-gonic/gin"

var r = gin.Default()

func CreatedRouter() *gin.Engine {

	Register(r)
	return r
}
