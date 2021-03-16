package sort

// bubble 冒泡排序
type bubble struct {
	Array
}

func NewBubbleSort() Sorter {
	return &bubble{Array{Type: "bubbleSort"}}
}

func (b *bubble) Sort(s []int) {
	b.slice = s
	bubbleSort(b.slice)
}

func bubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s) - i - 1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}