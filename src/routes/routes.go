package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jrschmidtt/api-gin/controllers"
)

func HandleRequest(){
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.HealthCheck)
	r.Run(":5000")
}