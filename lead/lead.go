package lead

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"github.com/shahruk25/go-fiber-crm/database"
)

type Lead struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLead() {}

func GetLeads(c *fiber.Ctx) {
	db := database.GetConn()
	rows, err := db.Query(context.Background(), "Select * from lead limit 10")
	if err != nil {
		log.Fatal(err)
	}
	var leads []Lead
	for rows.Next() {
		var lead Lead
		err = rows.Scan(&lead.Id, &lead.Name, &lead.Company, &lead.Email, &lead.Phone)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		leads = append(leads, lead)
	}
	c.JSON(&leads)
	defer rows.Close()
}

func NewLead(c *fiber.Ctx) {

}

func DeleteLead(c *fiber.Ctx) {

}
