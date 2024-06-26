package main

import (
	"math/rand"
	"testing"

	"github.com/iamsumit/go-dsa/algo/search/bs"
	"github.com/iamsumit/go-dsa/algo/search/es"
	"github.com/iamsumit/go-dsa/algo/search/is"
	"github.com/iamsumit/go-dsa/algo/search/js"
	"github.com/iamsumit/go-dsa/algo/search/ls"
	"github.com/iamsumit/go-dsa/algo/search/ts"
)

func BenchmarkLinearSearchFirstElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ls.Search(list, 1)
	}
}

func BenchmarkLinearSearch100thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ls.Search(list, 100)
	}
}

func BenchmarkLinearSearch500thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ls.Search(list, 500)
	}
}

func BenchmarkLinearSearch1000thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ls.Search(list, 1000)
	}
}

func BenchmarkLinearSearch5000thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ls.Search(list, 5000)
	}
}

func BenchmarkLinearSearch10000thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ls.Search(list, 10000)
	}
}

func BenchmarkLinearSearch50000thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ls.Search(list, 50000)
	}
}

func BenchmarkLinearSearchLastElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ls.Search(list, 99999)
	}
}

func BenchmarkLinearSearchElementNotPresent(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ls.Search(list, 100000)
	}
}

func BenchmarkLinearSearchRandomOrder(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, rand.Intn(100000))
	}

	for i := 0; i < b.N; i++ {
		ls.Search(list, i)
	}
}

func BenchmarkBinarySearch1000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 1000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		bs.Search(list, i)
	}
}

func BenchmarkBinarySearch10000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 10000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		bs.Search(list, i)
	}
}

func BenchmarkBinarySearch100000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		bs.Search(list, i)
	}
}

func BenchmarkBinarySearch100000000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		bs.Search(list, i)
	}
}

func BenchmarkBinarySearchRandomIncrement(b *testing.B) {
	list := []int{}
	var i int
	for len(list) < 100000000 {
		list = append(list, i)
		i += rand.Intn(100)
	}

	midElement := list[len(list)/2]
	for i := 0; i < b.N; i++ {
		bs.Search(list, midElement)
	}
}

func BenchmarkBinarySearchElementNotPresent(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		bs.Search(list, 100000)
	}
}

func BenchmarkTernarySearch1000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 1000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ts.Search(list, i)
	}
}

func BenchmarkTernarySearch10000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 10000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ts.Search(list, i)
	}
}

func BenchmarkTernarySearch100000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ts.Search(list, i)
	}
}

func BenchmarkTernarySearch100000000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ts.Search(list, i)
	}
}

func BenchmarkTernarySearchRandomIncrement(b *testing.B) {
	list := []int{}
	var i int
	for len(list) < 100000000 {
		list = append(list, i)
		i += rand.Intn(100)
	}

	midElement := list[len(list)/2]
	for i := 0; i < b.N; i++ {
		ts.Search(list, midElement)
	}
}

func BenchmarkTernarySearchElementNotPresent(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		ts.Search(list, 100000)
	}
}

func BenchmarkJumpSearchFirstElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		js.Search(list, 1)
	}
}

func BenchmarkJumpSearch100thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		js.Search(list, 100)
	}
}

func BenchmarkJumpSearch500thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		js.Search(list, 500)
	}
}

func BenchmarkJumpSearch1000thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		js.Search(list, 1000)
	}
}

func BenchmarkJumpSearch5000thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		js.Search(list, 5000)
	}
}

func BenchmarkJumpSearch10000thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		js.Search(list, 10000)
	}
}

func BenchmarkJumpSearch50000thElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		js.Search(list, 50000)
	}
}

func BenchmarkJumpSearchLastElement(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		js.Search(list, 99999)
	}
}

func BenchmarkJumpSearchElementNotPresent(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		js.Search(list, 100000)
	}
}

func BenchmarkJumpSearchElementNotPresentNegative(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		js.Search(list, -200)
	}
}

func BenchmarkInterpolationSearch1000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 1000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		is.Search(list, i)
	}
}

func BenchmarkInterpolationSearch10000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 10000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		is.Search(list, i)
	}
}

func BenchmarkInterpolationSearch100000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		is.Search(list, i)
	}
}

func BenchmarkInterpolationSearch100000000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		is.Search(list, i)
	}
}

func BenchmarkInterpolationSearchMidElement1(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		is.Search(list, 4502938)
	}
}

func BenchmarkInterpolationSearchMidElement2(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		is.Search(list, 45029380)
	}
}

func BenchmarkInterpolationSearchRandomIncrement(b *testing.B) {
	list := []int{}
	var i int
	for len(list) < 100000000 {
		list = append(list, i)
		i += rand.Intn(100)
	}

	midElement := list[len(list)/2]
	for i := 0; i < b.N; i++ {
		is.Search(list, midElement)
	}
}

func BenchmarkInterpolationSearchElementNotPresent(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		is.Search(list, 100000)
	}
}

func BenchmarkExponentialSearch1000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 1000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		es.Search(list, i)
	}
}

func BenchmarkExponentialSearch10000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 10000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		es.Search(list, i)
	}
}

func BenchmarkExponentialSearch100000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		es.Search(list, i)
	}
}

func BenchmarkExponentialSearch100000000Elements(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		es.Search(list, i)
	}
}

func BenchmarkExponentialSearchRandomIncrement(b *testing.B) {
	list := []int{}
	var i int
	for len(list) < 100000000 {
		list = append(list, i)
		i += rand.Intn(100)
	}

	midElement := list[len(list)/2]
	for i := 0; i < b.N; i++ {
		es.Search(list, midElement)
	}
}

func BenchmarkExponentialSearchElementNotPresent(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		es.Search(list, 100000)
	}
}

func BenchmarkExponentialSearchElementNotPresentNegative(b *testing.B) {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	for i := 0; i < b.N; i++ {
		es.Search(list, -200)
	}
}
