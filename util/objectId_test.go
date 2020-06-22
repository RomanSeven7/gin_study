package util

import "testing"

func BenchmarkNewObjectID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewObjectID()
	}
}
