package Keygen

import "testing"

func TestUuidGenerator_GetRandomKey(t *testing.T) {
	kg := NewUuidGenerator()
	storage := TestStorage{}
	key := kg.GetRandomKey(storage)
	if len(key) != 36 {
		t.Errorf("expected key to have 36 character, got : %v", key)
	}
}

