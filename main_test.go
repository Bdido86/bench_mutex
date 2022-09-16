// go test -bench=. -benchmem -benchtime=1000x
package main

import "testing"

func BenchmarkRunChannelExample(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		runChannelExample()
	}
}

func BenchmarkRunMutexExample(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		runMutexExample()
	}
}

func BenchmarkRunAtomicExample(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		runAtomicExample()
	}
}
