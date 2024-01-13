package main

import (
	"log"

	"github.com/CAndresFernandez/go-api/config"
	"github.com/CAndresFernandez/go-api/controllers"
	"github.com/CAndresFernandez/go-api/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := config.Connect()
	if err != nil {
		log.Fatal("failed to connect to the DB", err)
	}
	defer db.Close()

	db.AutoMigrate()

	handler := controllers.NewHandler(db)

	r := gin.New()
	routes.BookstoreRoutes(r, handler)
	r.Run()
}








