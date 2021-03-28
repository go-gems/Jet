package API

import (
	"errors"
	"jet/API/Keygen"
	"jet/StorageEngines"
	"testing"
)

type TestStorage struct {
	Key               string
	Content           []byte
	ShouldReturnError bool
}

func (t *TestStorage) KeyExists(id string) (bool, error) {
	if t.ShouldReturnError {
		return false, errors.New("something bad appended")
	}
	return id == t.Key, nil

}

func NewTestStorage(key string, content string) *TestStorage {
	return &TestStorage{Key: key, Content: []byte(content), ShouldReturnError: false}

}

func (t *TestStorage) Store(id string, content []byte) error {
	if t.ShouldReturnError {
		return errors.New("Error on Store")
	}
	t.Key = id
	t.Content = content
	return nil
}

func (t *TestStorage) FetchOne(id string) ([]byte, error) {
	if t.ShouldReturnError {
		return nil, errors.New("Error on FetchOne")
	}
	if t.Key != id {
		return nil, errors.New("unknown key")
	}
	return t.Content, nil
}

func (t *TestStorage) FetchAll() (map[string][]byte, error) {
	if t.ShouldReturnError {
		return nil, errors.New("Error on FetchAll")
	}
	return map[string][]byte{t.Key: t.Content}, nil
}

func (t *TestStorage) Delete(id string) error {
	if t.ShouldReturnError {
		return errors.New("Error on Delete")
	}
	if len(t.Key) < 1 {
		return errors.New("unknown key")
	}
	t.Key = ""
	t.Content = nil
	return nil
}

var _ StorageEngines.Storage = (*TestStorage)(nil)

func TestDelete(t *testing.T) {
	testStorageInstance := NewTestStorage("testKey", "Hello!")
	err := Delete(testStorageInstance, "testKey")
	if err != nil {
		t.Errorf("Expected error to be nil : got %v", err.Error())
	}

}


func TestStore(t *testing.T) {
	testStorageInstance := NewTestStorage("", "")

	got, err := Store(testStorageInstance,Keygen.UuidGenerator{}, "bonjour")
	if got != testStorageInstance.Key {
		t.Errorf("Expected key to be %v; expected %v", testStorageInstance.Key, got)
	}
	if err != nil {
		t.Errorf("Expected error to be nil : got %v", err.Error())
	}
}
func TestGet(t *testing.T) {
	testStorageInstance := NewTestStorage("key", "val")
	got, err := Get(testStorageInstance, "key")
	if got != "val" {
		t.Errorf("Expected key to be %v; Got %v", "val", got)
	}
	if err != nil {
		t.Errorf("Expected error to be nil : got %v", err.Error())
	}

	got2, err := Get(testStorageInstance, "empty-key")
	if got2 != "" {
		t.Errorf("Expected value to be empty; Got %v", got2)
	}
	if err == nil {
		t.Errorf("Expected error to not being nil")
	}
}

func TestAll(t *testing.T) {
	testStorageInstance := NewTestStorage("key", "val")
	content, err := All(testStorageInstance)
	if err != nil {
		t.Errorf("An Error has occured on fetching all data: %v", err.Error())
	}
	if content["key"] != "val" {
		t.Errorf("Content as key 'key' should be 'val', got %v", content["key"])
	}

	testStorageInstance.ShouldReturnError = true
	_, err2 := All(testStorageInstance)
	if err2 == nil {
		t.Errorf("An Error was expected")
	}

}
func TestAllKeys(t *testing.T) {
	testStorageInstance := NewTestStorage("key", "val")
	content, err := AllKeys(testStorageInstance)
	if err != nil {
		t.Errorf("An Error has occured on fetching all data: %v", err.Error())
	}
	if len(content) != 1 {
		t.Errorf("expected to have 1 content, got %v", len(content))
	}
	testStorageInstance.ShouldReturnError = true
	_, err2 := AllKeys(testStorageInstance)
	if err2 == nil {
		t.Errorf("An Error was expected")
	}

}
func TestSet(t *testing.T) {
	testStorageInstance := NewTestStorage("key", "val")
	key, err := Set(testStorageInstance, "key", "val2")
	if err != nil {
		t.Errorf("An Error has occured on fetching all data: %v", err.Error())
	}
	if key != "key" {
		t.Errorf("Content key should be 'key', got %v", key)
	}
	if string(testStorageInstance.Content) != "val2" {
		t.Errorf("invalid update, should be %v, got %v", "val2", string(testStorageInstance.Content))
	}

	testStorageInstance2 := NewTestStorage("", "")
	_, err2 := Set(testStorageInstance2, "", "val2")
	if err2 == nil {
		t.Errorf("An Error should have occured")
	}
}
func TestExists(t *testing.T) {
	testStorageInstance := NewTestStorage("key", "val")
	exists, err := Exists(testStorageInstance, "key")
	if !exists {
		t.Errorf("Key not found")
	}
	if err != nil {
		t.Errorf("An error occured: %v", err.Error())
	}

	exists, err = Exists(testStorageInstance, "key2")
	if exists {
		t.Errorf("Key found")
	}
	if err != nil {
		t.Errorf("An error occured: %v", err.Error())
	}
	testStorageInstance.ShouldReturnError = true
	exists, err = Exists(testStorageInstance, "key")
	if exists {
		t.Errorf("Key not found found")
	}
	if err == nil {
		t.Errorf("An error was expected")
	}
}
