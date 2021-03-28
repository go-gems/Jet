package main

import (
	"flag"
	"jet/API/Keygen"
	"jet/StorageEngines"
	"jet/http"
)

var storage StorageEngines.Storage = StorageEngines.NewInMemoryStorage()
var keyGenerator Keygen.KeyGenerator = Keygen.NewUuidGenerator()

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
	http.Serve(storage, keyGenerator)
}
