package StorageEngines

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func setUp() *FileStorage {
	os.Mkdir("test", 0777)
	ioutil.WriteFile("test/dev", []byte("hello"), 0777)
	return NewFileStorage("test")
}
func tearDown() {
	os.RemoveAll("./test")
}

func TestFileStorage_Store(t *testing.T) {
	fs := setUp()
	content := "dev2-content"
	err := fs.Store("dev2", []byte(content))
	if err != nil {
		t.Errorf("An error has occured on storage : %v", err.Error())
	}
	if _, err := os.Stat("test/dev2"); errors.Is(err, os.ErrNotExist) {
		t.Error("file test/dev2 does not exist")
	} else {
		fileContent, _ := ioutil.ReadFile("test/dev2")
		if string(fileContent) != content {
			t.Errorf("File content not match, expected %v, got %v", content, fileContent)
		}
	}
	tearDown()
}

func TestFileStorage_Delete(t *testing.T) {
	fs := setUp()
	fs.Delete("dev")
	if _, err := os.Stat("test/dev"); !errors.Is(err, os.ErrNotExist) {
		t.Error("file test/dev is not deleted")
	}
	tearDown()
}

func TestFileStorage_FetchOne(t *testing.T) {
	fs := setUp()
	content, err := fs.FetchOne("dev")
	if string(content) != "hello" || err != nil {
		t.Errorf("storage content at key 'dev' should be 'hello'; got %v with error %v", string(content), err.Error())
	}

	content2, err2 := fs.FetchOne("dev2")
	if len(content2) > 0 || err2 == nil {
		t.Errorf("Expected no response and error, got %v", string(content2))
	}
	tearDown()
}
func TestFileStorage_FetchAll(t *testing.T) {
	fs := setUp()
	content, err := fs.FetchAll()
	if err != nil {
		t.Errorf("Fetch was not expecting error, but received %v", err)
	}
	if len(content) != 1 {
		t.Errorf("Expecting 1 element in response, received %v", len(content))
	}
	value, ok := content["dev"]
	if !ok {
		t.Error("Expected to find 'dev' Key in content")
	} else if string(value) != "hello" {
		t.Errorf("Expected to find 'hello' for key, found %v", string(value))

	}
	os.Chmod("test/dev", 0000)
	_, err2 := fs.FetchAll()
	log.Println(err2)
	if err2 == nil {
		t.Error("Expected to get error but everything's fine")
	}
	tearDown()
	_, err3 := fs.FetchAll()
	if err3 == nil {
		t.Error("Expected to get error but everything's fine")
	}

}

func TestFileStorage_KeyExists(t *testing.T) {
	fs := setUp()

	if found, err := fs.KeyExists("dev");!found || err != nil {
		t.Errorf("The key was not found or an error occured: %v", err)
	}
	if found, err := fs.KeyExists("dev2");found || err != nil {
		t.Errorf("The key was not found or an error occured: %v", err)
	}

	tearDown()
}
