package storage

//Add service interface
type Storage interface {
	Save(string) (string, error)
	Close() error
}
