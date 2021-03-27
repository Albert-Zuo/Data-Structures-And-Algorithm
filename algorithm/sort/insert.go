package sort

func init() {
	RegisterSorter("insertSort", &insert{})
}

// 插入排序
type insert struct {}


func (i *insert) Sort(s []int) {
	insertSort(s)
}

func insertSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s[j] < s[j - 1]; j-- {
			swap(s, j, j - 1)
		}
	}
}
