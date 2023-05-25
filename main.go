package main

import (
	"log"

	m "github.com/berrylradianh/go_berryl-radian-hamesha/middlewares"
	r "github.com/berrylradianh/go_berryl-radian-hamesha/routes"
)

// func init() {
// 	u.InitMigrate()
// 	b.InitMigrate()
// 	bl.InitMigrate()
// }

func main() {
	log.Println("Starting the application...")
	e := r.InitRoutes()
	m.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8080"))
}
