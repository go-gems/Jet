package Keygen

import (
	"jet/StorageEngines"
	"math/rand"
	"time"
)

var _ KeyGenerator = (*ShortGenerator)(nil)


type ShortGenerator struct {
	baseSize int
}

func NewShortGenerator(baseSize int) *ShortGenerator {
	return &ShortGenerator{baseSize: baseSize}
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"

func (s *ShortGenerator) GetRandomKey(storage StorageEngines.Storage) string {
	rand.Seed(time.Now().UnixNano())
	length := s.baseSize
	var randy = ""
	for {
		for i := 0; i < length; i++ {
			u := alphabet[rand.Intn(len(alphabet))]
			randy += string(u)
		}
		exist, _ := storage.KeyExists(randy)
		if !exist {
			break
		}
		length++
		randy = ""
	}

	return randy
}