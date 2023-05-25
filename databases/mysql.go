package databases

import (
	"fmt"
	"log"

	ebl "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/blogs"
	eb "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/books"
	eu "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/users"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := "root"
	password := "secret-password"
	host := "34.30.158.49"
	port := "3306"
	name := "compute_service_alterra"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		name,
	)

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		&eu.User{},
		&eb.Book{},
		&ebl.Blog{},
	)
}
