package storage

type Storage interface {
	Open() error
	Close() error
	FileStorage
}

type FileStorage interface {
	Create(file File) error
	Get(osFileName string, inode uint64) []File
	GetAll() []File
	Delete(osFileName string)
}
