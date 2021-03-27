package sort

func init() {
	RegisterSorter("shellSort", &shell{})
}

// 希尔排序
type shell struct {}



func(a *shell) Sort(s []int) {
	shellSort(s)
}

func shellSort(s []int) {
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