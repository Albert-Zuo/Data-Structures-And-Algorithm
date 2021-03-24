package sort

// bubble 冒泡排序
type arrayBubble struct {
	Array
}

func NewArrayBubble() ArraySorter {
	return &arrayBubble{Array{Type: "bubbleSort"}}
}

func (b *arrayBubble) Sort(s []int) {
	b.slice = s
	arrayBubbleSort(b.slice)
}

func arrayBubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s) - i - 1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}