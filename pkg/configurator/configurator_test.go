package configurator

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	data := `{
  "host": "localhost",
  "port": 5432,
  "user": "postgres",
  "password": "postgres",
  "dbname": "test1",
  "schedule": "@every 10s",
  "reports_path": "/Users/julialagun/GolandProjects/CVSReportGenerator/data"
}
`
	byteArrayData := []byte(data)
	var cfg Config
	if err := cfg.parseConfig(byteArrayData); err != nil {
		t.Fatalf("cannot parse config: %v", err)
	}

	testCfg := Config{Host: "localhost", Port: 5432, User: "postgres", Password: "postgres", Dbname: "test1",
		Schedule: "@every 10s", ReportsPath: "/Users/julialagun/GolandProjects/CVSReportGenerator/data"}

	if cfg != testCfg {
		t.Errorf("expected:\n%v\ngot:\n%v", testCfg, cfg)
	}
}
