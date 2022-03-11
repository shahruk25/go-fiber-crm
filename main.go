package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"github.com/jackc/pgx/v4"
	"github.com/shahruk25/go-fiber-crm/database"
	"github.com/shahruk25/go-fiber-crm/lead"
)

var db *pgx.Conn

func main() {
	app := fiber.New()
	initDB()
	setUpRoutes(app)
	app.Listen(3000)
	defer db.Close(context.Background())
}

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/newLead", lead.NewLead)
	app.Delete("/api/v1/delLead/:id", lead.DeleteLead)
}

func initDB() {
	database.GetConn()
	db = database.GetConn()
	log.Print("conn : ", db.Config())
}
