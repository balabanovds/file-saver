package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/balabanovds/file-saver/pkg/ticker"
)

var (
	fromDir  string
	toDir    string
	interval int
	dbFile   string
)

func init() {
	flag.StringVar(&fromDir, "from", "", "watch directory (mandatory)")
	flag.StringVar(&toDir, "to", "", "destination directory (mandatory)")
	flag.StringVar(&dbFile, "db", "./db/db.sqlite", "file for sqlite db")
	flag.IntVar(&interval, "interval", 3000, "polling interval in seconds")
}

func main() {
	flag.Parse()
	if fromDir == "" && toDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	done := ticker.Every(time.Duration(interval)*time.Second, processFiles)

	select {
	case <-sigCh:
		done <- true
	case <-done:
	}

}

func processFiles(t time.Time) error {
	println("processFiles called")
	return nil
}
