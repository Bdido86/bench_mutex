// go test -bench=. -benchmem -benchtime=1000x
package main

import "testing"

func BenchmarkRunAtomicExample(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		runAtomicExample()
	}
}

func BenchmarkRunMutexExample(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		runMutexExample()
	}
}

func BenchmarkRunMutexErrorExample(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		runMutexErrorExample()
	}
}

func BenchmarkRunChannelExample(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		runChannelExample()
	}
}
