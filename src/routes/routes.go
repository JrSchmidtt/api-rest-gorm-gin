package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jrschmidtt/api-gin/controllers"
)

func HandleRequest(){
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", controllers.RenderIndex)
	r.GET("/alunos", controllers.GetAllAluno)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.FindAlunoById)
	r.GET("/alunos/cpf/:cpf", controllers.FindAlunoByCPF)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.NoRoute(controllers.RotaNaoEncontrada)

	r.Run(":5000")
}