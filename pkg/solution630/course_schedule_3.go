package solution630

import (
	"container/heap"
	"fmt"
	"sort"
)

//There are n different online courses numbered from 1 to n.
//You are given an array courses where courses[i] = [durationi, lastDayi] indicate that
//the ith course should be taken continuously for durationi days and must be finished before or on lastDayi.
//
//You will start on the 1st day and you cannot take two or more courses simultaneously.
//Input: courses = [[100,200],[200,1300],[1000,1250],[2000,3200]]
//Output: 3
//Input: courses = [[1,2]]
//Output: 1
//Input: courses = [[3,2],[4,3]]
//Output: 0
func scheduleCourse(courses [][]int) int {
	//sort by lastDay of course
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	fmt.Printf("sorted by last day: %v \n", courses)

	h := &Heap{}
	total := 0 // total time used for courses
	for _, course := range courses {
		//if we can take this course, push it's duration into heap,
		//otherwise check and replace the longer course in heap
		if t := course[0]; total+t <= course[1] {
			total += t
			heap.Push(h, t)
			fmt.Printf("take the course: %v, now in heap: %v \n", course, h)
		} else if h.Len() > 0 && t < h.IntSlice[0] {
			total += t - h.IntSlice[0]
			h.IntSlice[0] = t
			heap.Fix(h, 0)
			fmt.Printf("replace the course in heap with: %v \n", course)
		}
	}
	return h.Len()
}

//Heap implement sort.Interface and Heap interface
type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return x
}
