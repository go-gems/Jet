// +build !skip

package main

import (
	"flag"
	"fmt"
	"jet/API/Keygen"
	"jet/StorageEngines"
	"jet/http"
	"log"
)

var storage StorageEngines.Storage = StorageEngines.NewInMemoryStorage()
var keyGenerator Keygen.KeyGenerator = Keygen.NewUuidGenerator()
var host = flag.String("host", "127.0.0.1", "Host to use")
var port = flag.String("port", "8000", "Port to use")

func init() {
	storageDir := flag.String("storage.dir", "", "Specify directory to use Storage Directory")
	generatorShort := flag.Int("generator.short", 0, "Specify the number of characters to start with short generation")
	generatorIncrement := flag.Bool("generator.inc", false, "Use increment to generate keys")
	flag.Parse()

	if len(*storageDir) > 0 {
		storage = StorageEngines.NewFileStorage(*storageDir)
	}

	if *generatorIncrement {
		keyGenerator = Keygen.NewIncrementGenerator()
	}
	if *generatorShort > 0 {
		keyGenerator = Keygen.NewShortGenerator(*generatorShort)
	}
}

func main() {

	hostAndPort := fmt.Sprintf("%v:%v",*host,*port)

	log.Fatal(http.Serve( hostAndPort , storage, keyGenerator))
}
