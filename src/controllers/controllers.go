package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrschmidtt/api-gin/database"
	"github.com/jrschmidtt/api-gin/models"
)

func GetAllAluno(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}