package main

import (
	"math/rand"
	"testing"

	"github.com/iamsumit/go-dsa/algo/insertionsort/is"
)

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