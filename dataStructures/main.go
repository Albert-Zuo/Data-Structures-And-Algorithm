package main

import (
	"fmt"
	"github.com/Albert-Zuo/Data-Structures-And-Algorithm/dataStructures/array"
)


func main() {

	var tests = []struct {
		value string
		index int
	}{
		{"a", 0},
		{"b", 1},
		{"c", 2},
		{"d", 3},

		{"e", 4},
		{"f", 5},
	}
	list := array.NewArrayByCapacity(-6)

	for _, test := range tests {
		list.Add(test.index, test.value)
	}
	fmt.Println(list)
	fmt.Println("==============================")

	iterator := list.Iterator();
	for iterator.HasNext() {
		next, _ := iterator.Next()
		if next == "e" {
			iterator.Remove()
		}
		fmt.Println(next)
	}
	fmt.Println("==============================")
	fmt.Println(list)

}
