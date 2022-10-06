package main

import (
	"fmt"
	"log"

	"github.com/This-Is-Prince/go-fiber-crm-basic/database"
	"github.com/This-Is-Prince/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error

	database.DB, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: Failed to connect database: %v", err)
	}
	log.Println("Connection opened to database...")

	err = database.DB.AutoMigrate(&lead.Lead{})
	if err != nil {
		log.Fatalf("Error: Failed to run auto migrate for given models: %v", err)
	}
	log.Println("Auto Migration successfully...")
}

func main() {
	fmt.Println("GO Fiber CRM Basic")

	initDatabase()
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}
