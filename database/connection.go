package database

import (
	"fmt" 
	"strconv"

	"github.com/kgermando/mspos-api/models"
	"github.com/kgermando/mspos-api/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// const DNSs =  "root:12345678@tcp(localhost:3306)/msposdb?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() {
	p := utils.Env("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}

	DNS := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.Env("DB_USER"), utils.Env("DB_PASSWORD"), utils.Env("DB_HOST"), port, utils.Env("DB_NAME"))
	connection, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	fmt.Println("Database Connected!")

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
