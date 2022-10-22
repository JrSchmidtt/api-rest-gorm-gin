package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jrschmidtt/api-gin/controllers"
)

func HandleRequest(){
	r := gin.Default()
	r.GET("/alunos", controllers.GetAllAluno)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.FindAlunoById)
	r.Run(":5000")
}