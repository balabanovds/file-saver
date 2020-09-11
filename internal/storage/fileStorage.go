package storage

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3" //nolint:golint
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
	if err == nil {
		log.Printf("db connected successfully to %s\n", f.filename)
	}
	return
}

func (f *fileStorage) Close() error {
	log.Println("closing db")
	return f.db.Close()
}

func (f *fileStorage) Create(file File) error {
	_, err := f.db.NamedExec("insert into files (os_name, app_name, inode) "+
		"values (:os_name, :app_name, :inode)", map[string]interface{}{
		"os_name":  file.OSName,
		"app_name": file.AppName,
		"inode":    file.Inode,
	})
	if err != nil {
		return err
	}
	return nil
}

func (f *fileStorage) Count(osFileName string, inode uint64) int {
	var result int

	err := f.db.Get(&result, "select count(*) from files where os_name = $1 and inode = $2", osFileName, inode)
	if err != nil {
		log.Printf("db: error %v\n", err)
	}
	return result
}

func (f *fileStorage) Delete(osFileName string) {
	_, _ = f.db.Exec("delete from files where os_name = $1", osFileName)
}
