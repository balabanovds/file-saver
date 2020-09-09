package storage

type File struct {
	ID          int
	OSName      string    `db:"os_name"`
	AppName     string    `db:"app_name"`
	Inode       uint64    `db:"inode"`
}

func NewFile(osName, appName string, inode uint64) File {
	return File{
		OSName:      osName,
		AppName:     appName,
		Inode:       inode,
	}
}

func (f File) Match(o File) bool {
	return f.OSName == o.OSName && f.Inode == o.Inode
}
