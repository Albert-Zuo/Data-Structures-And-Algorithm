package main

import (
	"fmt"
	"github.com/Albert-Zuo/Data-Structures-And-Algorithm/algorithm/sort"
)

var (
	sortList = []sort.Sorter {
		sort.NewBubbleSort(),
		sort.NewQuickSort(),
	}
)

func main() {

	for _, sorter := range sortList {
		slice := []int{1, 2, 5, 3, 2, 1, -20, 299, 33, 45}
		sorter.Sort(slice)
		fmt.Println(sorter)
	}

}

