package stack

type Stack interface {
	GetSize() int       // 获取栈的长度
	IsEmpty() bool      // 栈是否为空
	Push(e interface{}) // 进栈
	Pop() interface{}   // 出栈
	Peek() interface{}  // 查看栈顶元素
}
