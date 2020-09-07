package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type fileStorage struct {
	filename string
	db       *sqlx.DB
}

func New(filename string) Storage {
	return &fileStorage{filename: filename}
}

func (f *fileStorage) Open() (err error) {
	f.db, err = sqlx.Connect("sqlite3", f.filename)
	return
}

func (f *fileStorage) Close() error {
	return f.db.Close()
}

func (f *fileStorage) Create(file *File) error {
	_, err := f.db.NamedQuery("insert into file (os_name, app_name, size, os_created_at) "+
		"values (:os_name, :app_name, :size, :os_created_at)", map[string]interface{}{
		"os_name":       file.OSName,
		"app_name":      file.AppName,
		"size":          file.Size,
		"os_created_at": file.CreatedTime,
	})
	if err != nil {
		return err
	}
	return nil
}

func (f *fileStorage) Get(osName string) []File {
	result := []File{}

	err := f.db.Get(&result, "select * from files where os_name = $1", osName)
	if err != nil {
		log.Printf("db: error %v\n", err)
	}
	return result
}

func (f *fileStorage) GetAll() []File {
	result := []File{}

	err := f.db.Get(&result, "select * from files")
	if err != nil {
		log.Printf("db: error %v\n", err)
	}
	return result
}
