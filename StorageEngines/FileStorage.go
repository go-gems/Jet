package StorageEngines

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var _ Storage = (*FileStorage)(nil)

type FileStorage struct {
	directory string
}

func (f FileStorage) fullKey(id string) string {
	return fmt.Sprintf("%v/%v", f.directory, id)

}
func (f *FileStorage) Store(id string, content []byte) error {
	var err error = nil

	err = ioutil.WriteFile(f.fullKey(id), content, 0644)
	return err

}

func (f *FileStorage) FetchOne(id string) ([]byte, error) {
	var err error = nil
	b, err := ioutil.ReadFile(f.fullKey(id))
	return b, err
}

func (f *FileStorage) FetchAll() (map[string][]byte, error) {
	response := map[string][]byte{}
	err := filepath.Walk(f.directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		content, err := ioutil.ReadFile(path)
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		response[info.Name()] = content
		return err
	})
	return response, err
}
func (f *FileStorage) Delete(id string) error {
	err := os.Remove(f.fullKey(id))
	return err

}

func (f *FileStorage) KeyExists(id string) (bool, error) {
	_, err := os.Stat(f.fullKey(id))

	if err != nil {

		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil

}

func NewFileStorage(directory string) *FileStorage {
	return &FileStorage{directory: directory}
}
