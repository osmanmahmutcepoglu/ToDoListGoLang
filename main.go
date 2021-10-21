package main

import (
	 "ToDoListGoLang/Controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	var router = gin.Default()
	router.Use(gin.Logger()) //Debug modu
	router.Use(cors())

	toDoControllers := controllers.NewTodoController(controllers.NewBaseController())
	v1 := router.Group("api/v1")
	{
		v1.GET("getTodoList", toDoControllers.GetTodoList)
		v1.POST("addTodo", toDoControllers.AddTodo)
	}
	router.Run("localhost:8080")
}
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}