package Keygen

import "testing"

func TestShortGenerator_GetRandomKey(t *testing.T) {
	storage := TestStorage{}
	for i := 1; i < 128; i++ {
		kg := NewShortGenerator(i)
		key := kg.GetRandomKey(storage)
		if len(key) != i {
			t.Errorf("expected key to have %v char, got %v", i, len(key))
		}
	}
	storage.sizeUnavailable =10
	kg := NewShortGenerator(1)
	key := kg.GetRandomKey(storage)
	if len(key) != 11 {
		t.Errorf("expected key to have %v char, got %v", 11, len(key))


	}
}