package API

import (
	"errors"
	"jet/API/Keygen"
	"jet/StorageEngines"
)

func Store(s StorageEngines.Storage, keyGenerator Keygen.KeyGenerator, item string) (string, error) {
	uid := keyGenerator.GetRandomKey(s)
	err := s.Store(uid, []byte(item))
	return uid, err
}
func Delete(s StorageEngines.Storage, id string) error {
	return s.Delete(id)
}

func Set(s StorageEngines.Storage, id string, item string) (string, error) {
	if len(id) < 1 {
		return "", errors.New("cannot Set")
	}
	err := s.Store(id, []byte(item))
	return id, err
}

func Get(s StorageEngines.Storage, id string) (string, error) {
	content, err := s.FetchOne(id)
	return string(content), err
}

func All(s StorageEngines.Storage) (map[string]string, error) {
	content := map[string]string{}
	res, err := s.FetchAll()
	if err != nil {
		return nil, err
	}
	for k, v := range res {
		content[k] = string(v)
	}
	return content, nil
}

func AllKeys(s StorageEngines.Storage) ([]string, error) {
	content := []string{}
	res, err := s.FetchAll()
	if err != nil {
		return nil, err
	}
	for k, _ := range res {
		content = append(content, k)
	}
	return content, nil
}

func Exists(s StorageEngines.Storage, key string) (bool, error) {
	return s.KeyExists(key)
}
