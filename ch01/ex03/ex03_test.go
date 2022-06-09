package main

import (
	"strings"
	"testing"
)

var args = []string{"Go", "with", "the", "flow", "is", "apparently", "a", "song"}

func Concatenation(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
}

func BenchmarkConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Concatenation(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}
