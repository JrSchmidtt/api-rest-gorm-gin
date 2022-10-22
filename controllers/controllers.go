package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jrschmidtt/api-gin/models"
)

func ExibeTodosAlunos(c *gin.Context){
	c.JSON(200, models.Alunos)
}

func HealthCheck(c *gin.Context){
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:":"E ai " + nome + ", eu estou ONLINE !",
	})
}