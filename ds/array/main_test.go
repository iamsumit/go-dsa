package main

import (
	"testing"
)

func BenchmarkArrayAccessAll(b *testing.B) {
	numSeries := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = numSeries
	}
}

func BenchmarkArrayAccessOne(b *testing.B) {
	numSeries := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = numSeries[1]
	}
}

func BenchmarkArrayAccessSelected(b *testing.B) {
	numSeries := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = numSeries[1:3]
	}
}

func BenchmarkArrayUpdate(b *testing.B) {
	numSeries := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		numSeries[0] = 10
		numSeries[1] = 20
		numSeries[2] = 30
		numSeries[3] = 40
		numSeries[4] = 50
		_ = numSeries
	}
}

func BenchmarkArrayIteration(b *testing.B) {
	numSeries := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		for i, num := range numSeries {
			_ = i
			_ = num
		}
	}
}

func BenchmarkFindFirstElementInArray(b *testing.B) {
	numSeries := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 10 {
				break
			}
		}
	}
}

func BenchmarkFindThirdElementInArray(b *testing.B) {
	numSeries := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 30 {
				break
			}
		}
	}
}

func BenchmarkFindLastElementInArray(b *testing.B) {
	numSeries := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 50 {
				break
			}
		}
	}
}

func BenchmarkMatrixAccessAll(b *testing.B) {
	matrix := [2][2]int{{11, 12}, {21, 22}}
	for i := 0; i < b.N; i++ {
		_ = matrix
	}
}

func BenchmarkMatrixAccessOne(b *testing.B) {
	matrix := [2][2]int{{11, 12}, {21, 22}}
	for i := 0; i < b.N; i++ {
		_ = matrix[1]
	}
}

func BenchmarkMatrixAccessSelected(b *testing.B) {
	matrix := [2][2]int{{11, 12}, {21, 22}}
	for i := 0; i < b.N; i++ {
		_ = matrix[1][1]
	}
}

func BenchmarkMatrixUpdate(b *testing.B) {
	matrix := [2][2]int{{11, 12}, {21, 22}}
	for i := 0; i < b.N; i++ {
		matrix[0][0] = 10
		matrix[0][1] = 20
		matrix[1][0] = 30
		matrix[1][1] = 40
		_ = matrix
	}
}

func BenchmarkMatrixIteration(b *testing.B) {
	matrix := [2][2]int{{11, 12}, {21, 22}}
	for i := 0; i < b.N; i++ {
		for i, row := range matrix {
			for j, num := range row {
				_ = i
				_ = j
				_ = num
			}
		}
	}
}
