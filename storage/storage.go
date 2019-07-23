package storage

//Service storage interface
type Service interface {
	Save(string) (string, error)
	Close() error
}
