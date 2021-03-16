package array

type List interface {
	Size() int // 获取 List 长度

	Get(index int) (interface{}, error) // 根据索引获取 element

	Set(index int, e interface{}) // 修改索引对应的值

	Add(index int, e interface{}) // 在索引位置添加新的 element

	Append(e interface{}) // 末尾添加 element

	Remove(index int) interface{} // 根据索引删除 element

	Clean() // 清理数组内存

	Iterator() Iterator // 获取迭代器
}
