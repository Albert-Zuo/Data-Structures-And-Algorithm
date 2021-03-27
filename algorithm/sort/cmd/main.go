package main

import (
	"fmt"
	"github.com/Albert-Zuo/Data-Structures-And-Algorithm/algorithm/sort"
)

// test
func main() {
	var (
		count = len(sort.KeySet)
		done  = make(chan struct{}, count)
	)
	// 修改为 并发执行打印函数
	sort.MgrF(func(currType string, sortedSlice []int) {
		go func() {
			fmt.Println(fmt.Sprintf("%s: %v", currType, sortedSlice))
			done <- struct{}{}
		}()
	})
	for name := range sort.KeySet {
		s := []int{1, 2, 5, 3, 2, 1, -20, 299, 33, 45}
		err := sort.DoSort(name, s)
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < count; i++ {
		<- done
	}

}
