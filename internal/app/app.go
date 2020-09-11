package app

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/balabanovds/file-saver/internal/storage"
)

const (
	layoutISO = "2006-01-02"
)

type App struct {
	st       storage.Storage
	fromDir  string
	toDir    string
	interval time.Duration
}

func New(st storage.Storage, fromDir, toDir string, seconds int) *App {
	return &App{
		st:       st,
		fromDir:  fromDir,
		toDir:    toDir,
		interval: time.Duration(seconds) * time.Second,
	}
}

func (a *App) Run(ctx context.Context) {
	every(ctx, a.interval, a.processFiles)
}

func (a *App) processFiles(t time.Time) error {
	return filepath.Walk(a.fromDir, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsDir() {
			return nil
		}
		if !info.Mode().IsRegular() {
			log.Printf("file %s is not regular. skipping..", path)
			return nil
		}

		statT, ok := info.Sys().(*syscall.Stat_t)
		if !ok {
			log.Printf("filewalk: failed to get inode for file %s", path)
			return nil
		}

		appFileName := t.Format(layoutISO) + "_" + info.Name()
		file := storage.NewFile(info.Name(), appFileName, statT.Ino)

		filesFound := a.st.Count(info.Name(), statT.Ino)
		if filesFound != 0 {
			log.Printf("file %s already in DB\n", info.Name())
			return nil
		}

		log.Printf("found new file %s; ino %d;\n", info.Name(), statT.Ino)
		if err := a.st.Create(file); err != nil {
			return err
		}

		fromAbs, err := filepath.Abs(a.fromDir)
		if err != nil {
			return err
		}

		toAbs, err := filepath.Abs(a.toDir)
		if err != nil {
			return err
		}

		from := filepath.Join(fromAbs, info.Name())
		to := filepath.Join(toAbs, appFileName)

		if err := copyFile(from, to); err != nil {
			log.Printf("failed to copy file: %v\n", err)
			a.st.Delete(info.Name())
			return err
		}
		log.Printf("file %s copied to %s\n", from, to)

		return nil
	})
}

func copyFile(from, to string) error {
	src, err := os.Open(from)
	if err != nil {
		return err
	}

	dst, err := os.Create(to)
	if err != nil {
		return err
	}
	defer func() {
		if err := src.Close(); err != nil {
			log.Printf("failed to close src file %s\n", from)
		}
		if err := dst.Close(); err != nil {
			log.Printf("failed to close dst file %s\n", to)
		}
	}()
	_, err = io.Copy(dst, src)
	return err
}
