package runner

import (
	"database/sql"
	"log"
	"os"
	"path"
	"time"

	"github.com/robfig/cron/v3"

	"csv-reports-generator/pkg/configurator"
	"csv-reports-generator/pkg/dbreader"
	reportgen "csv-reports-generator/pkg/report"
)

type Runner struct {
	scheduler *cron.Cron
	db        *sql.DB
}

func NewRunner() *Runner {
	return &Runner{
		scheduler: cron.New(cron.WithSeconds()),
	}
}

func (r *Runner) Run(fileName string) {
	cfg := configurator.LoadConfig(fileName)
	db, err := dbreader.OpenDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	r.db = db
	err = r.runSchedule(cfg, r.scheduleFunc(cfg.ReportsPath, db))
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Runner) scheduleFunc(path string, db *sql.DB) func() {
	return func() {
		if err := makeReport(path, db); err != nil {
			log.Println("error generating report:", err)
		}
	}
}

func (r *Runner) Stop() {
	r.scheduler.Stop()
	r.db.Close()
}

func (r *Runner) runSchedule(cfg configurator.Config, fn func()) error {
	if _, err := r.scheduler.AddFunc(cfg.Schedule, fn); err != nil {
		return err
	}
	r.scheduler.Start()
	return nil
}

// genFileName generates report file name based on current time
func genFileName(outputDirectory string) string {
	return path.Join(outputDirectory, time.Now().Format(time.DateTime)+".csv")
}

// makeReport reads data from database and writes reports to the outputDirectory
func makeReport(outputDirectory string, db *sql.DB) error {
	rows, err := dbreader.GetData(db)
	if err != nil {
		return err
	}
	reportFileName := genFileName(outputDirectory)
	f, err := os.Create(reportFileName)
	if err != nil {
		return err
	}
	defer f.Close()
	return reportgen.CreateCSVFile(f, rows)
}
