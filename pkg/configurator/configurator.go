package configurator

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host        string `json:"host"`
	Port        uint16 `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Dbname      string `json:"dbname"`
	Schedule    string `json:"schedule"`
	ReportsPath string `json:"reports_path"`
}

func (c *Config) parseConfig(data []byte) error {
	return json.Unmarshal(data, &c)
}

func LoadConfig(fileName string) Config {
	cfg := Config{
		Host:        os.Getenv("CSV_DBHOST"),
		User:        os.Getenv("CSV_DBUSER"),
		Password:    os.Getenv("CSV_DBPASS"),
		Dbname:      os.Getenv("CSV_DBNAME"),
		Schedule:    os.Getenv("CSV_SCHEDULE"),
		ReportsPath: os.Getenv("CSV_REPORTS_PATH"),
	}
	if portStr := os.Getenv("CSV_DBPORT"); portStr != "" {
		if port, err := strconv.ParseUint(portStr, 10, 16); err != nil {
			log.Print("CSV_DBPORT is defined but not a number")
		} else {
			cfg.Port = uint16(port)
		}
	}
	if fileName != "" {
		data, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalf("Error openning config file: %v", err)
		}
		if err := cfg.parseConfig(data); err != nil {
			log.Fatalf("Error parsing config file: %v", err)
		}
	}
	if cfg.Host == "" {
		log.Fatalf("Database host is missing")
	}
	if cfg.Port == 0 {
		log.Fatalf("Database port is missing")
	}
	if cfg.User == "" {
		log.Fatalf("DataBase user name is missing")
	}
	if cfg.Password == "" {
		log.Fatalf("DataBase password is missing")
	}
	if cfg.Dbname == "" {
		log.Fatalf("DataBase name is missing")
	}
	if cfg.Schedule == "" {
		log.Fatalf("Reports schedule is missing")
	}
	if cfg.ReportsPath == "" {
		log.Fatalf("Reports path is missing")
	}
	return cfg
}
