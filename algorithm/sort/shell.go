package sort


// 快排
type arrayShell struct {
	Array
}

func NewArrayShell() ArraySorter {
	return &arrayShell{Array{Type: "shellSort"}}
}

func(a *arrayShell) Sort(s []int) {
	a.slice = s
	arrayShellSort(a.slice[:])
}

func arrayShellSort(s []int) {
	n := len(s)
	h := 1
	for h < n/3 { //寻找合适的间隔h
		h = 3*h + 1
	}
	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h && s[j] < s[j-h]; j -= h {
				swap(s, j, j-h)
			}
		}
		h /= 3
	}
}