package StorageEngines

import "testing"

func TestInMemoryStorage_Store(t *testing.T) {
	storage := NewInMemoryStorage()
	storage.Store("key", []byte("val"))
	value, ok := storage.content["key"]
	if !ok {
		t.Errorf("Key 'key' should be stored")
	}
	if string(value) != "val" {
		t.Errorf("Value for key should be val, got %v", string(value))
	}
}
func TestInMemoryStorage_FetchOne(t *testing.T) {
	storage := NewInMemoryStorage()
	storage.content["key"] = []byte("value")
	result, err := storage.FetchOne("key")
	if string(result) != "value" {
		t.Errorf("FetchOne was expecting 'value', got %v", string(result))
	}
	if err != nil {
		t.Errorf("FetchOne was not expecting error, error '%v' encountered", err.Error())
	}

	result2, err2 := storage.FetchOne("inexistantKey")
	if string(result2) != "" {
		t.Errorf("FetchOne was not expecting content, got %v", string(result2))
	}
	if err2 == nil {
		t.Error("FetchOne was expecting error, but everything worked strangely fine")
	}
}

func TestInMemoryStorage_Delete(t *testing.T) {
	storage := NewInMemoryStorage()
	storage.content["key"] = []byte("value")
	err := storage.Delete("key")
	if len(storage.content) > 0 {
		t.Errorf("The inMemory storage should be empty, data found : %v", storage.content)
	}
	if err != nil {
		t.Errorf("Delete was not expecting error, but received %v", err.Error())
	}
	err = storage.Delete("not-existing-key")
	if err == nil {
		t.Error("Delete was expecting error, but nothing wrong happened")
	}
}

func TestInMemoryStorage_FetchAll(t *testing.T) {
	storage := NewInMemoryStorage()
	storage.content["key"] = []byte("value")
	content, err := storage.FetchAll()
	if err != nil {
		t.Errorf("Fetch was not expecting error, but received %v", err)
	}
	if len(content) != 1 {
		t.Errorf("Expecting 1 element in response, received %v", len(content))
	}
	value, ok := content["key"]
	if !ok {
		t.Error("Expected to find 'key' Key in content")
	} else if string(value) != "value" {
		t.Errorf("Expected to find 'value' for key, found %v", string(value))

	}
}

func TestInMemoryStorage_KeyExists(t *testing.T) {
	storage := NewInMemoryStorage()
	storage.content["key"] = []byte("value")
	if res, err := storage.KeyExists("key"); !res || err != nil {
		t.Error("Key not found nor an error occured")
	}
	if res, err := storage.KeyExists("key2"); res || err != nil {
		t.Error("Key found or something went wrong")
	}

}