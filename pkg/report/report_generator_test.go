package report

import (
	"bytes"
	"testing"
)

func TestCsvGeneration(t *testing.T) {
	customers := []Customer{
		{
			Id:           1,
			Name:         "White Rabbit",
			EmailAddress: "white.rabbit@animal.com",
		},
		{
			Id:           2,
			Name:         "Fire Fox",
			EmailAddress: "firefox@animal.com",
		},
	}
	out := new(bytes.Buffer)
	err := CreateCSVFile(out, customers)
	if err != nil {
		t.Fatal(err)
	}
	expected :=
		`id, name, email_address
1,White Rabbit,white.rabbit@animal.com
2,Fire Fox,firefox@animal.com
`
	if out.String() != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, out.String())
	}
}
