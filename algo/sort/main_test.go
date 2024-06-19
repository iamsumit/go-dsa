package main

import (
	"math/rand"
	"testing"

	"github.com/iamsumit/go-dsa/algo/sort/bs"
	"github.com/iamsumit/go-dsa/algo/sort/hs"
	"github.com/iamsumit/go-dsa/algo/sort/is"
	"github.com/iamsumit/go-dsa/algo/sort/ms"
	"github.com/iamsumit/go-dsa/algo/sort/qs"
	"github.com/iamsumit/go-dsa/algo/sort/ss"
)

func BenchmarkBubbleSort10Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		bs.Sort(newArr)
	}
}

func BenchmarkBubbleSort100Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(100))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		bs.Sort(newArr)
	}
}

func BenchmarkBubbleSort1000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(1000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		bs.Sort(newArr)
	}
}

func BenchmarkBubbleSort10000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(10000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		bs.Sort(newArr)
	}
}

func BenchmarkBubbleSort100000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		bs.Sort(newArr)
	}
}

func BenchmarkBubbleSort100000ElementsAsc(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		bs.Sort(newArr)
	}
}

func BenchmarkBubbleSort100000ElementsDesc(b *testing.B) {
	arr := []int{}
	for i := 100000; i > 0; i-- {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		bs.Sort(newArr)
	}
}

func BenchmarkSelectionSort10Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ss.Sort(newArr)
	}
}

func BenchmarkSelectionSort100Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(100))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ss.Sort(newArr)
	}
}

func BenchmarkSelectionSort1000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(1000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ss.Sort(newArr)
	}
}

func BenchmarkSelectionSort10000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(10000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ss.Sort(newArr)
	}
}

func BenchmarkSelectionSort100000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ss.Sort(newArr)
	}
}

func BenchmarkSelectionSort100000ElementsAsc(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ss.Sort(newArr)
	}
}

func BenchmarkSelectionSort100000ElementsDesc(b *testing.B) {
	arr := []int{}
	for i := 100000; i > 0; i-- {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ss.Sort(newArr)
	}
}

func BenchmarkInsertionSort10Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		is.Sort(newArr)
	}
}

func BenchmarkInsertionSort100Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(100))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		is.Sort(newArr)
	}
}

func BenchmarkInsertionSort1000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(1000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		is.Sort(newArr)
	}
}

func BenchmarkInsertionSort10000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(10000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		is.Sort(newArr)
	}
}

func BenchmarkInsertionSort100000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		is.Sort(newArr)
	}
}

func BenchmarkInsertionSort100000ElementsAsc(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		is.Sort(newArr)
	}
}

func BenchmarkInsertionSort100000ElementsDesc(b *testing.B) {
	arr := []int{}
	for i := 100000; i > 0; i-- {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		is.Sort(newArr)
	}
}

func BenchmarkQuickSort10Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		qs.Sort(newArr)
	}
}

func BenchmarkQuickSort100Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(100))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		qs.Sort(newArr)
	}
}

func BenchmarkQuickSort1000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(1000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		qs.Sort(newArr)
	}
}

func BenchmarkQuickSort10000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(10000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		qs.Sort(newArr)
	}
}

func BenchmarkQuickSort100000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		qs.Sort(newArr)
	}
}

func BenchmarkQuickSort100000ElementsAsc(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		qs.Sort(newArr)
	}
}

func BenchmarkQuickSort100000ElementsDesc(b *testing.B) {
	arr := []int{}
	for i := 100000; i > 0; i-- {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		qs.Sort(newArr)
	}
}

func BenchmarkMergeSort10Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ms.Sort(newArr)
	}
}

func BenchmarkMergeSort100Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(100))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ms.Sort(newArr)
	}
}

func BenchmarkMergeSort1000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(1000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ms.Sort(newArr)
	}
}

func BenchmarkMergeSort10000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(10000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ms.Sort(newArr)
	}
}

func BenchmarkMergeSort100000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ms.Sort(newArr)
	}
}

func BenchmarkMergeSort100000ElementsAsc(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ms.Sort(newArr)
	}
}

func BenchmarkMergeSort100000ElementsDesc(b *testing.B) {
	arr := []int{}
	for i := 100000; i > 0; i-- {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		ms.Sort(newArr)
	}
}

func BenchmarkHeapSort10Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		hs.Sort(newArr)
	}
}

func BenchmarkHeapSort100Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(100))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		hs.Sort(newArr)
	}
}

func BenchmarkHeapSort1000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(1000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		hs.Sort(newArr)
	}
}

func BenchmarkHeapSort10000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(10000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		hs.Sort(newArr)
	}
}

func BenchmarkHeapSort100000Elements(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100000))
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		hs.Sort(newArr)
	}
}

func BenchmarkHeapSort100000ElementsAsc(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		hs.Sort(newArr)
	}
}

func BenchmarkHeapSort100000ElementsDesc(b *testing.B) {
	arr := []int{}
	for i := 100000; i > 0; i-- {
		arr = append(arr, i)
	}

	for i := 0; i < b.N; i++ {
		// Copy the original array to avoid sorting the already sorted array.
		newArr := append([]int(nil), arr...)
		hs.Sort(newArr)
	}
}
