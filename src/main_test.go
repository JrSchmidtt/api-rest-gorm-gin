package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jrschmidtt/api-gin/controllers"
	"github.com/jrschmidtt/api-gin/database"
	"github.com/jrschmidtt/api-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock(){
	aluno := models.Aluno{Nome: "Nome do aluno Teste", 
	CPF: "12345678901",
	RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock(){
	var aluno models.Aluno
	database.DB.Delete(&aluno)
	ID = int(aluno.ID)
}

func TestGetAllAlunoHandler(t *testing.T) {
	database.ConnectPostgresDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.GetAllAluno)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestFindAlunoByCPFHandler(t *testing.T){
	database.ConnectPostgresDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.FindAlunoByCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestFindAlunoByIDHandler(t *testing.T){
	database.ConnectPostgresDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.FindAlunoById)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var alunoMock models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Nome)
	assert.Equal(t, "Nome do aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
}

func TestDeleteAlunoHandler(t *testing.T){
	database.ConnectPostgresDB()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t,http.StatusOK, res.Code)
}

func TestEditaAlunoHandler(t *testing.T){
	database.ConnectPostgresDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome do aluno Teste", 
	CPF: "12345670000",
	RG: "123456789"}
	varJson, _ := json.Marshal(aluno)
	pathToEdit := "/aluno/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathToEdit, bytes.NewBuffer(varJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var newAlunoMock models.Aluno
	err := json.Unmarshal(res.Body.Bytes(), &newAlunoMock)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "12345670000", newAlunoMock.CPF)
	assert.Equal(t, "12345670000", newAlunoMock.RG)
}