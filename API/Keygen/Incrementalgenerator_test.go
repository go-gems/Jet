package Keygen

import (
	"strconv"
	"testing"
)

func TestIncrementGenerator_GetRandomKey(t *testing.T) {
	increment := NewIncrementGenerator()
	storage := TestStorage{}
	for i := 1; i <= 10; i++ {
		val := increment.GetRandomKey(storage)
		if val != strconv.Itoa(i) {
			t.Errorf("Expected to have %v, got %v", i, val)
		}
	}
	increment2 := NewIncrementGenerator()
	storage.sizeUnavailable = 2
	val2 := increment2.GetRandomKey(storage)
	if val2 != "100" {
		t.Errorf("Expected to have %v, got %v", "100", val2)
	}
}
