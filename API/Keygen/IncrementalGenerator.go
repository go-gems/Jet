package Keygen

import (
	"jet/StorageEngines"
	"strconv"
	"sync"
)

type IncrementGenerator int

var incrementMutex = sync.Mutex{}

func NewIncrementGenerator() *IncrementGenerator {
	var ig IncrementGenerator = 0
	return &ig
}

func (i *IncrementGenerator) GetRandomKey(storage StorageEngines.Storage) string {
	incrementMutex.Lock()
	defer incrementMutex.Unlock()
	str := ""
	for {
		*i++
		str = strconv.Itoa(int(*i))
		ok, _ := storage.KeyExists(str)
		if !ok {
			break
		}
	}

	return str
}
