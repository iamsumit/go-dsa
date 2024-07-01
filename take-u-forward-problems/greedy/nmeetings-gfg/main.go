package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/n-meetings-in-one-room/
// Using sorting for better calculation
func maximumMeetings(start []int, end []int) int {
	if len(start) <= 1 {
		return 1
	}

	mcount := len(start)
	meetings := []int{}
	for i := 0; i < mcount; i++ {
		meetings = append(meetings, i)
	}

	slices.SortFunc(meetings, func(i, k int) int {
		if end[i] < end[k] {
			return -1
		} else if end[i] > end[k] {
			return 1
		} else {
			return 0
		}
	})

	counter := 1
	e := end[meetings[0]]
	for i := 1; i < mcount; i++ {
		if start[meetings[i]] > e {
			e = end[meetings[i]]
			counter++
		}
	}

	return counter
}

// func maximumMeetings(start []int, end []int) []int {
// 	mcount := len(start)
// 	meetings := []int{1}
// 	e := end[0]
// 	for i := 1; i < mcount; i++ {
// 		if start[i] > e {
// 			e = end[i]
// 			meetings = append(meetings, i+1)
// 		}
// 	}

// 	return meetings
// }

func main() {
	fmt.Println(maximumMeetings([]int{1, 3, 0, 5, 8, 5}, []int{2, 4, 5, 7, 9, 9}))
	fmt.Println(maximumMeetings([]int{1, 3, 0, 5, 9, 7}, []int{2, 4, 5, 6, 10, 8}))
	fmt.Println(maximumMeetings([]int{1, 3, 0, 5, 9, 7}, []int{2, 7, 5, 6, 10, 8}))
	fmt.Println(maximumMeetings([]int{75250, 50074, 43659, 8931, 11273, 27545, 50879, 77924}, []int{112960, 114515, 81825, 93424, 54316, 35533, 73383, 160252}))
}
