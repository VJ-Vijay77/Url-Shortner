package main

import (
	"log"
	"urlShortner/controller"
	"urlShortner/database"
	"urlShortner/model"
	"urlShortner/router"
	"urlShortner/shared"
)

func main() {
	db, err := database.ConnectDb()
	if err != nil {
		log.Fatal("Cant connect to db", err)
	}
	db.AutoMigrate(model.Url{})
	userService := shared.NewUserService(db)
	userHandler := controller.NewUserHandler(userService)
	r := router.Router(userHandler)
	r.Run() 
}
