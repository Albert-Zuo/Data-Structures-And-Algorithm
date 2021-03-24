package sort

type arrayInsert struct {
	Array
}

func NewArrayInsert() ArraySorter {
	return &arrayInsert{Array{Type: "insertSort"}}
}

func (i *arrayInsert) Sort(s []int) {
	i.slice = s
	arrayBubbleSort(s)
}

func arrayInsertSort(s []int) {
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
