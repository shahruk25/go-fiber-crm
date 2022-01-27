package lead

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"github.com/shahruk25/go-fiber-crm/database"
)

type Lead struct {
	id      int
	name    string
	company string
	email   string
	phone   int
}

func GetLead() {}

func GetLeads(c *fiber.Ctx) {
	db := database.GetDB()
	rows, err := db.Query(context.Background(), "Select * from lead limit 10")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
	}

}

func NewLead(c *fiber.Ctx) {

}

func DeleteLead(c *fiber.Ctx) {

}
