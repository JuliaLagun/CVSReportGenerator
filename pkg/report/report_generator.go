package report

import (
	"fmt"
	"io"
)

type Customer struct {
	Id           int
	Name         string
	EmailAddress string
}

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
