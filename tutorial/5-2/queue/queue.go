// 让Queue支持任何类型
package queue

type Queue []interface{}

// 任意类型
/*func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}*/

// 限定只能Push进队列 int 类型
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// 未在函数形参中定义限定的类型，这样只能在运行时才能发现错误
// 不推荐这样做
/*func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int))
}*/

// 任意类型
/*func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}*/

// 限定只能Pop出 int 类型
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

// 未在函数返回值处定义限定的类型，这样只能在运行时才能发现错误
// 不推荐这样做
/*func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}*/

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
