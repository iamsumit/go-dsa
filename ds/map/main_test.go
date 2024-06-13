package main

import "testing"

func BenchmarkMapInsert(b *testing.B) {
	mapVar := map[int]int{}
	for i := 0; i < b.N; i++ {
		mapVar[i] = i
	}
}

func BenchmarkMapInsertWithPreallocation(b *testing.B) {
	mapVar := make(map[int]int, b.N)
	for i := 0; i < b.N; i++ {
		mapVar[i] = i
	}
}

func BenchmarkMapUpdate(b *testing.B) {
	mapVar := map[int]int{}
	for i := 0; i < 10000; i++ {
		mapVar[i] = i
	}

	for i := 0; i < b.N; i++ {
		mapVar[i] = i + 1
	}
}

func BenchmarkMapDelete(b *testing.B) {
	mapVar := map[int]int{}
	for i := 0; i < 10000; i++ {
		mapVar[i] = i
	}

	for i := 0; i < b.N; i++ {
		delete(mapVar, i)
	}
}

func BenchmarkAccessMap(b *testing.B) {
	mapVar := map[int]int{}
	for i := 0; i < 10000; i++ {
		mapVar[i] = i
	}

	for i := 0; i < b.N; i++ {
		_, ok := mapVar[i]
		_ = ok
	}
}

func BenchmarkMapIterationOver100(b *testing.B) {
	mapVar := map[int]int{}
	for i := 0; i < 100; i++ {
		mapVar[i] = i
	}

	for i := 0; i < b.N; i++ {
		for key, value := range mapVar {
			_ = key
			_ = value
		}
	}
}

func BenchmarkMapIterationOver1000(b *testing.B) {
	mapVar := map[int]int{}
	for i := 0; i < 1000; i++ {
		mapVar[i] = i
	}

	for i := 0; i < b.N; i++ {
		for key, value := range mapVar {
			_ = key
			_ = value
		}
	}
}

func BenchmarkMapIterationOver10000(b *testing.B) {
	mapVar := map[int]int{}
	for i := 0; i < 10000; i++ {
		mapVar[i] = i
	}

	for i := 0; i < b.N; i++ {
		for key, value := range mapVar {
			_ = key
			_ = value
		}
	}
}
