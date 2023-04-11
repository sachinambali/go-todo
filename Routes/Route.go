package Routes

import (
	"go-todo-app/Controllers"
	"go-todo-app/Middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
		v1.POST("/token", Controllers.GenerateToken)
		v1.POST("/user/register", Controllers.RegisterUser)
		secured := v1.Group("/secured").Use(Middlewares.Auth())
		{
			secured.GET("/ping", Controllers.Ping)
		}
	}
	return r
}
