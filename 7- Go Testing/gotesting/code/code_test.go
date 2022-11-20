package code

import (
	"testing"
)

// go test ./...
// go test ./...
func TestIsValied(t *testing.T) {
	nID, err := NewNatinalCode("3356947567")

	if err != nil {
		t.Error(err)
	}

	if nID.IsValid() != true {
		t.Errorf("Problem in validation")
	}
}

// go test -bench=.
func BenchmarkIsValied(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewNatinalCode("3356947567")

		if err != nil {
			b.Error(err)
		}

	}
}
