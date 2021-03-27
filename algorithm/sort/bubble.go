package sort

func init() {
	RegisterSorter("bubbleSort", &bubble{})
}

// bubble 冒泡排序
type bubble struct {}


func (b *bubble) Sort(s []int) {
	bubbleSort(s)
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