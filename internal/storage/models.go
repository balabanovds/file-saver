package storage

import "time"

type File struct {
	ID          int
	OSName      string    `db:"os_name"`
	AppName     string    `db:"app_name"`
	Size        int64     `db:"size"`
	CreatedTime time.Time `db:"os_created_at"`
}
