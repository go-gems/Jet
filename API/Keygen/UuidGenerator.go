package Keygen

import (
	"github.com/google/uuid"
	"jet/StorageEngines"
)

// Generates a Uuid as key
type UuidGenerator struct{}
var _ KeyGenerator = (*UuidGenerator)(nil)

func NewUuidGenerator() *UuidGenerator {
	return &UuidGenerator{}
}

func (u UuidGenerator) GetRandomKey(storage StorageEngines.Storage) string {
	return uuid.New().String()
}
