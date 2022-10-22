package controllers

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(c *gin.Context){
	c.JSON(200, gin.H{
		"id":"1",
		"nome": "Junior Schmidt",
	})
}

func HealthCheck(c *gin.Context){
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:":"E ai " + nome + ", eu estou ONLINE !",
	})
}