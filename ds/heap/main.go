package main

import (
	"fmt"

	"github.com/iamsumit/go-dsa/ds/heap/h"
)

func main() {
	fmt.Println("------ Heap Data Structure ------")
	hs := h.Heap{}
	fmt.Println("Value of hs: ", hs)

	hs.Insert(1)
	hs.Insert(2)
	hs.Insert(3)

	fmt.Println("--------------------")

	fmt.Println("Peek after inserting 1,2,3:", hs.Peek())

	fmt.Println("--------------------")

	hs.Insert(0)

	fmt.Println("Peek after inserting 0:", hs.Peek())

	fmt.Println("--------------------")

	hs.Insert(4)
	hs.Insert(5)
	hs.Insert(6)
	hs.Insert(7)
	hs.Insert(8)
	hs.Insert(9)
	hs.Insert(10)
	hs.Insert(11)
	hs.Insert(12)
	hs.Insert(13)
	hs.Insert(14)
	fmt.Println("Insert 4-14 in the order:", hs)

	fmt.Println("--------------------")

	hs.Delete(10)
	fmt.Println("Delete 10:", hs)

	fmt.Println("--------------------")

	hs.Delete(13)
	fmt.Println("Delete 13:", hs)

	fmt.Println("--------------------")

	hs.Delete(4)
	fmt.Println("Delete 4:", hs)

	hs.Delete(6)
	fmt.Println("Delete 6:", hs)

	fmt.Println("--------------------")

	hs.UpdateVal(0, 15)
	fmt.Println("Update index 0 to 15:", hs)

	fmt.Println("--------------------")

	hs.UpdateVal(2, 8)
	fmt.Println("Update index 2 to 8:", hs)

	fmt.Println("--------------------")

	i := hs.FindIndex(9)
	hs.IncreaseByValue(i, 4)
	fmt.Printf("Increase 9 by 4 at index %d: %v\n", i, hs)

	fmt.Println("--------------------")

	i = hs.FindIndex(13)
	hs.DecreaseByValue(i, 4)
	fmt.Printf("Decrease 13 by 4 at index %d: %v\n", i, hs)
}
