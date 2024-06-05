package main

import (
	"strings"
	"testing"
)

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := "Hello, World!"
		_ = message
	}
}

func BenchmarkStringIndex(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = message[0]
	}
}

func BenchmarkStringSlice(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = message[0:5]
	}
}

func BenchmarkStringConcatenation(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = message + " Welcome to Go!"
	}
}

func BenchmarkStringConcatenationWithStringsPkg(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{message, "Welcome to Go!"}, " ")
	}
}

func BenchmarkStringIterating(b *testing.B) {
	message := "Hello!⌘"
	for i := 0; i < b.N; i++ {
		for i, char := range message {
			_ = i
			_ = char
		}
	}
}

func BenchmarkStringContains(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.Contains(message, "World")
	}
}

func BenchmarkStringSearch(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.Index(message, "World")
	}
}

func BenchmarkStringPrefix(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.HasPrefix(message, "Hello")
	}
}

func BenchmarkStringSuffix(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.HasSuffix(message, "World!")
	}
}

func BenchmarkStringReplace(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.Replace(message, "World", "Go", -1)
	}
}

func BenchmarkStringReplaceAll(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.ReplaceAll(message, "World", "Go")
	}
}

func BenchmarkStringEqual(b *testing.B) {
	message1 := "Hello, World!"
	message2 := "Hello, world!"
	for i := 0; i < b.N; i++ {
		_ = message1 == message2
	}
}

func BenchmarkStringEqualFold(b *testing.B) {
	message1 := "Hello, World!"
	message2 := "Hello, world!"
	for i := 0; i < b.N; i++ {
		_ = strings.EqualFold(message1, message2)
	}
}

func BenchmarkStringSplit(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.Split(message, ", ")
	}
}

func BenchmarkStringTrim(b *testing.B) {
	message := " Hello, World! "
	for i := 0; i < b.N; i++ {
		_ = strings.TrimSpace(message)
	}
}

func BenchmarkStringTrimPrefix(b *testing.B) {
	message := "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.TrimPrefix(message, "Hello")
	}
}
