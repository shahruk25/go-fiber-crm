package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"github.com/jackc/pgx/v4"
	"github.com/shahruk25/go-fiber-crm/database"
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
	// app.Get(GetLeads)
	// app.Get(GetLead)
	// app.Post(NewLead)
	// app.Delete(DeleteLead)

}

func initDB() {
	database.GetConn()
	db = database.GetDB()
	log.Print("conn : ", db.Config())
}
