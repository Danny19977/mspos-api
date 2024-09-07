package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/kgermando/mspos-api/models"
	"github.com/kgermando/mspos-api/utils"
	// "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBSQL *sql.DB

func Connect() {
	p := utils.Env("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}

	DNS := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", utils.Env("DB_HOST"), port, utils.Env("DB_USER"), utils.Env("DB_PASSWORD"), utils.Env("DB_NAME"))
	connection, err := gorm.Open(postgres.Open(DNS), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("Could not connect to the database 😵😰")
	}

	// DNS := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.Env("DB_USER"), utils.Env("DB_PASSWORD"), utils.Env("DB_HOST"), port, utils.Env("DB_NAME"))
	// connection, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	// if err != nil {
	// 	panic("could not connect to the database")
	// }

	DB = connection
	fmt.Println("Database Connected!")

	 
	dbSQL, err := sql.Open("postgres", DNS)
    if err != nil {
        panic(err)
    }
	DBSQL = dbSQL
   
	fmt.Println("Database Connected SQL!")

	connection.AutoMigrate(
		&models.User{},
		&models.Province{},
		&models.Area{},
		&models.Asm{},
		&models.Manager{},
		&models.Pos{},
		&models.PosForm{},
		&models.UserLogs{},
		&models.Sup{},
	)

}
