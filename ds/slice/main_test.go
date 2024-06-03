package main

import "testing"

func BenchmarkSlicecAccessEmpty(b *testing.B) {
	numSeries := make([]int, 2, 5)
	for i := 0; i < b.N; i++ {
		_ = numSeries
	}
}

func BenchmarkSliceAccessAll(b *testing.B) {
	numSeries := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = numSeries
	}
}

func BenchmarkSliceAccessOne(b *testing.B) {
	numSeries := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = numSeries[1]
	}
}

func BenchmarkSliceAccessSelected(b *testing.B) {
	numSeries := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = numSeries[1:3]
	}
}

func BenchmarkSliceUpdate(b *testing.B) {
	numSeries := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		numSeries[0] = 10
		numSeries[1] = 20
		numSeries[2] = 30
		numSeries[3] = 40
		numSeries[4] = 50
		_ = numSeries
	}
}

func BenchmarkSliceAppend(b *testing.B) {
	numSeries := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		numSeries = append(numSeries, 6)
		_ = numSeries
	}
}

func BenchmarkSliceIteration(b *testing.B) {
	numSeries := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		for i, num := range numSeries {
			_ = i
			_ = num
		}
	}
}

func BenchmarkFindFirstElementInSlice(b *testing.B) {
	numSeries := []int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 10 {
				break
			}
		}
	}
}

func BenchmarkFindThirdElementInSlice(b *testing.B) {
	numSeries := []int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 30 {
				break
			}
		}
	}
}

func BenchmarkFindLastElementInSlice(b *testing.B) {
	numSeries := []int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 50 {
				break
			}
		}
	}
}
