package storage

type Storage interface {
	Open() error
	Close() error
	FileStorage
}

type FileStorage interface {
	Create(file *File) error
	GetAll() []File
}
