package report

import (
	"fmt"
	"io"
)

// Customer data structure represents database customer table row
type Customer struct {
	Id           int
	Name         string
	EmailAddress string
}

// CreateCSVFile writes an array of customer data to the output writer (`out`) in csv format
func CreateCSVFile(out io.Writer, customers []Customer) error {
	header := "id, name, email_address"
	if _, err := fmt.Fprintln(out, header); err != nil {
		return err
	}

	for _, client := range customers {
		if _, err := fmt.Fprintf(out, "%d,%s,%s\n", client.Id, client.Name, client.EmailAddress); err != nil {
			return err
		}
	}
	return nil
}
