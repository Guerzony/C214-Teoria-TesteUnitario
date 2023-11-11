package database

import (
	"C214-teoria-GO/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Conceta_BD() {
	sdc := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable "
	//String de conexao
	DB, err = gorm.Open(postgres.Open(sdc))
	if err != nil {
		log.Panic("erro ao conectar ao banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
