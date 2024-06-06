package main

import (
	"testing"
)

func BenchmarkArrayAccess5All(b *testing.B) {
	numSeries := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = numSeries
	}
}

func BenchmarkArrayAccess100All(b *testing.B) {
	numSeries := [100]int{}
	for i := 0; i < 100; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		_ = numSeries
	}
}

func BenchmarkArrayAccess1000All(b *testing.B) {
	numSeries := [1000]int{}
	for i := 0; i < 1000; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		_ = numSeries
	}
}

func BenchmarkArrayAccessOneIn5(b *testing.B) {
	numSeries := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = numSeries[1]
	}
}

func BenchmarkArrayAccessOneIn100(b *testing.B) {
	numSeries := [100]int{}
	for i := 0; i < 100; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		_ = numSeries[1]
	}
}

func BenchmarkArrayAccessOneIn1000(b *testing.B) {
	numSeries := [1000]int{}
	for i := 0; i < 1000; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		_ = numSeries[1]
	}
}

func BenchmarkArrayAccessSelectedIn5(b *testing.B) {
	numSeries := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = numSeries[1:3]
	}
}

func BenchmarkArrayAccessSelectedIn100(b *testing.B) {
	numSeries := [100]int{}
	for i := 0; i < 100; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		_ = numSeries[10:20]
	}
}

func BenchmarkArrayAccessSelectedIn1000(b *testing.B) {
	numSeries := [1000]int{}
	for i := 0; i < 1000; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		_ = numSeries[100:200]
	}
}

func BenchmarkArrayUpdateIn5(b *testing.B) {
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

func BenchmarkArrayUpdateIn100(b *testing.B) {
	numSeries := [100]int{}
	for i := 0; i < 100; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		numSeries[44] = 50
	}
}

func BenchmarkArrayUpdateIn1000(b *testing.B) {
	numSeries := [1000]int{}
	for i := 0; i < 1000; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		numSeries[44] = 50
	}
}

func BenchmarkArrayIterationIn5(b *testing.B) {
	numSeries := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		for i, num := range numSeries {
			_ = i
			_ = num
		}
	}
}

func BenchmarkArrayIterationIn100(b *testing.B) {
	numSeries := [100]int{}
	for i := 0; i < 100; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		for i, num := range numSeries {
			_ = i
			_ = num
		}
	}
}

func BenchmarkArrayIterationIn1000(b *testing.B) {
	numSeries := [1000]int{}
	for i := 0; i < 1000; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		for i, num := range numSeries {
			_ = i
			_ = num
		}
	}
}

func BenchmarkFindFirstElementInArrayIn5(b *testing.B) {
	numSeries := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 10 {
				break
			}
		}
	}
}

func BenchmarkFindFirstElementInArrayIn100(b *testing.B) {
	numSeries := [100]int{}
	for i := 0; i < 100; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 10 {
				break
			}
		}
	}
}

func BenchmarkFindFirstElementInArrayIn1000(b *testing.B) {
	numSeries := [1000]int{}
	for i := 0; i < 1000; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 10 {
				break
			}
		}
	}
}

func BenchmarkFindThirdElementInArrayIn5(b *testing.B) {
	numSeries := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 30 {
				break
			}
		}
	}
}

func BenchmarkFindThirtiethElementInArrayIn100(b *testing.B) {
	numSeries := [100]int{}
	for i := 0; i < 100; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 30 {
				break
			}
		}
	}
}

func BenchmarkFindThreeHundredthElementInArrayIn1000(b *testing.B) {
	numSeries := [1000]int{}
	for i := 0; i < 1000; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 300 {
				break
			}
		}
	}
}

func BenchmarkFindLastElementInArrayIn5(b *testing.B) {
	numSeries := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 50 {
				break
			}
		}
	}
}

func BenchmarkFindLastElementInArrayIn100(b *testing.B) {
	numSeries := [100]int{}
	for i := 0; i < 100; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 99 {
				break
			}
		}
	}
}

func BenchmarkFindLastElementInArrayIn1000(b *testing.B) {
	numSeries := [1000]int{}
	for i := 0; i < 1000; i++ {
		numSeries[i] = i
	}

	for i := 0; i < b.N; i++ {
		for _, num := range numSeries {
			if num == 999 {
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
