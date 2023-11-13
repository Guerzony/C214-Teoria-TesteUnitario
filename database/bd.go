package database

import (
	"C214-teoria-GO/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conecta_BD() DBInterface {
	sdc := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable "
	//String de conexao
	DB, err = gorm.Open(postgres.Open(sdc))
	if err != nil {
		log.Panic("erro ao conectar ao banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
	return &Database{DB: DB}
}

type DBInterface interface {
	Find(dest interface{}) DBInterface
	First(dest interface{}, conds ...interface{}) DBInterface
	Create(value interface{}) DBInterface
	Delete(value interface{}, conds ...interface{}) DBInterface
	Model(value interface{}) DBInterface
	UpdateColumns(value interface{}) DBInterface
	Where(value interface{}, args ...interface{}) DBInterface
}

type Database struct {
	*gorm.DB
}

func (db *Database) Find(dest interface{}) DBInterface {
	db.DB.Find(dest)
	return db
}

func (db *Database) First(dest interface{}, conds ...interface{}) DBInterface {
	db.DB.First(dest, conds...)
	return db
}

func (db *Database) Create(value interface{}) DBInterface {
	db.DB.Create(value)
	return db
}

func (db *Database) Delete(value interface{}, conds ...interface{}) DBInterface {
	db.DB.Delete(value, conds...)
	return db
}

func (db *Database) Model(value interface{}) DBInterface {
	db.DB.Model(value)
	return db
}

func (db *Database) UpdateColumns(value interface{}) DBInterface {
	db.DB.UpdateColumns(value)
	return db
}

func (db *Database) Where(value interface{}, args ...interface{}) DBInterface {
	db.DB.Where(value, args...)
	return db
}

var (
	DB  *gorm.DB
	err error
)
