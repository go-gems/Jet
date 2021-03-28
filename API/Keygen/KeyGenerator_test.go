package Keygen

type TestStorage struct {
	sizeUnavailable int
}

func (t TestStorage) Store(id string, content []byte) error { return nil }

func (t TestStorage) FetchOne(id string) ([]byte, error) { return []byte("ok"), nil }

func (t TestStorage) FetchAll() (map[string][]byte, error) {
	return map[string][]byte{"k": []byte("ok")}, nil
}

func (t TestStorage) Delete(id string) error { return nil }

func (t TestStorage) KeyExists(id string) (bool, error) { return len(id) <= t.sizeUnavailable, nil }

