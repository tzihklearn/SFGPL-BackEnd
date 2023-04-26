package route

import (
	"SFGPL-End/biz/handler"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	program := r.Group("/program")
	{
		program.GET("/all", handler.GetAllProgram)
		program.GET("/search", handler.Search)
		program.POST("/add", handler.Add)
		program.POST("/deleted", handler.Deleted)
		program.POST("/update", handler.Update)
	}

}
