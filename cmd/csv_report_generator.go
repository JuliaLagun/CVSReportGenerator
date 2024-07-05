package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"csv-reports-generator/pkg/runner"
)

// command line argument - path to config.json (optional)
func main() {
	r := runner.NewRunner()
	if len(os.Args) < 2 {
		r.Run("")
	} else {
		r.Run(os.Args[1])
	}
	defer r.Stop()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
