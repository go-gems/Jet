package Keygen

import (
	"jet/StorageEngines"
)

type KeyGenerator interface {
	GetRandomKey(storage StorageEngines.Storage) string
}

