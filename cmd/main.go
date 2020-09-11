package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/balabanovds/file-saver/internal/app"
	"github.com/balabanovds/file-saver/internal/storage"
	"golang.org/x/net/context"
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
	flag.IntVar(&interval, "interval", 3600, "polling interval in seconds")
}

func main() {
	flag.Parse()
	if fromDir == "" && toDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	dbFileAbs, err := filepath.Abs(dbFile)
	if err != nil {
		log.Fatalln(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	st := storage.New(dbFileAbs)
	defer func() {
		if err := st.Close(); err != nil {
			log.Printf("db close error: %v\n", err)
		}
	}()

	if err := st.Open(); err != nil {
		log.Fatalf("db open error: err")
	}

	if err := os.Mkdir(toDir, 0700); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Fatalln(err)
		}
	}

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	app.New(st, fromDir, toDir, interval).Run(ctx)

	<-sigCh
}
