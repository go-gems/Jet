package StorageEngines

type Storage interface {
	Store(id string, content []byte) error
	FetchOne(id string) ([]byte, error)
	FetchAll() (map[string][]byte, error)
	Delete(id string) error
	KeyExists(id string) (bool, error)
}