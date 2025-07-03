package main

import "testing"

func BenchmarkKthCharacter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kthCharacter(450)
	}
}

func BenchmarkKthCharacterEditorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kthCharacterEditorial(450)
	}
}
