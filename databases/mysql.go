package databases

import (
	"fmt"
	"log"
	"os"

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

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

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
