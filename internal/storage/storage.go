package storage

type Storage interface {
	Open() error
	Close() error
	FileStorage
}

type FileStorage interface {
	Create(file File) error
	Count(osFileName string, inode uint64) int
	Delete(osFileName string)
}
