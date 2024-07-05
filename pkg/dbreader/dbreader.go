package dbreader

// package for database related code
import (
	"database/sql"
	"fmt"
	"log"

	"csv-reports-generator/pkg/configurator"
	reportgen "csv-reports-generator/pkg/report"
)

func makeConnectionString(cfg configurator.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)
}

// OpenDB opens database using connection parameters from the config
func OpenDB(cfg configurator.Config) (*sql.DB, error) {
	postgresqlDbInfo := makeConnectionString(cfg)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// GetData returns array of Customers from the DB
func GetData(db *sql.DB) ([]reportgen.Customer, error) {
	rows, err := db.Query("SELECT id, name, email_address FROM customers")
	if err != nil {
		return nil, err
	}
	return transformData(rows), nil
}

func transformData(rows *sql.Rows) []reportgen.Customer {
	var customers []reportgen.Customer

	for rows.Next() {
		var client reportgen.Customer
		if err := rows.Scan(&client.Id, &client.Name, &client.EmailAddress); err != nil {
			log.Fatal(err)
		}
		customers = append(customers, client)
	}
	return customers
}
