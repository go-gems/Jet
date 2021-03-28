package Keygen

import (
	"jet/StorageEngines"
	"strconv"
)

type IncrementGenerator int

func NewIncrementGenerator() *IncrementGenerator {
	var ig IncrementGenerator = 0
	return &ig
}

func (i *IncrementGenerator) GetRandomKey(storage StorageEngines.Storage) string {
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
