package main

import "testing"

// Run with: go test -bench=. -benchmem ./04-arrays-slices-maps-loops

const benchSize = 100_000

func BenchmarkWithoutPrealloc(b *testing.B) {
	for b.Loop() {
		fillWithoutPrealloc(benchSize)
	}
}

func BenchmarkWithPrealloc(b *testing.B) {
	for b.Loop() {
		fillWithPrealloc(benchSize)
	}
}
