package config

import (
	"fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
    . "rest-api-go/service/models"
)

func DBInit() *gorm.DB {

	viper.SetConfigFile(`rest-api-go/service/config/config.json`)
	err := viper.ReadInConfig()

	if err != nil{
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Running on Debug Mode")
	}

	dbhost     := viper.GetString(`database.host`)
	port       := viper.GetString(`database.port`)
	username   := viper.GetString(`database.user`)
	password   := viper.GetString(`database.pass`)
	dbname 	   := viper.GetString(`database.name`)

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s", dbhost, port, username, dbname, password)
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	

	// add struct to add table in database
	conn.Debug().AutoMigrate(&User{})

	return conn
}
