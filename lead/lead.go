package lead

import (
	"context"
	"fmt"
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
	// IS_DEL  bool   `json:"isDel"`
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.GetConn()
	row, err := db.Query(context.Background(), "Select name,company,email,phone,id from lead where id=$1 and is_del =false", &id)
	defer row.Close()
	defer db.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if !(row.Next()) {
		c.JSON("No lead found")
		return
	}
	var lead Lead
	for row.Next() {
		err := row.Scan(&lead.Name, &lead.Company, &lead.Email, &lead.Phone, &lead.Id)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
	}
	c.JSON(lead)
}

func GetLeads(c *fiber.Ctx) {
	db := database.GetConn()
	rows, err := db.Query(context.Background(), "Select name,company,email,phone,id from lead where is_del =false")
	defer rows.Close()
	defer db.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var leads []Lead
	for rows.Next() {
		var lead Lead
		err = rows.Scan(&lead.Name, &lead.Company, &lead.Email, &lead.Phone, &lead.Id)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		leads = append(leads, lead)
	}
	c.JSON(&leads)
}

func NewLead(c *fiber.Ctx) {
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db := database.GetConn()
	defer db.Close(context.Background())
	sqlInsrt := `INSERT INTO lead (name,company,email,phone) VALUES ($1,$2,$3,$4) RETURNING id`
	var id int64
	err := db.QueryRow(c.Context(), sqlInsrt, lead.Name, lead.Company, lead.Email, lead.Phone).Scan(&id)

	if err != nil {

		log.Fatalf("unable to execute query. %v", err)
	}
	c.JSON(id)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.GetConn()
	res, err := db.Exec(context.Background(), "Update lead set is_del = true where id=$1", &id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	// check how many rows affected
	rowsAffected := res.RowsAffected()
	if rowsAffected < 1 {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	c.JSON(id)
}
