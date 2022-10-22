package database

import (
	"fmt"
	"log"

	"github.com/jrschmidtt/api-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectPostgresDB(){
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Fatal("Erro ao connectar com banco")
	}
	fmt.Println("Conectado com sucesso ao banco !")
	DB.AutoMigrate(models.Aluno{})
}