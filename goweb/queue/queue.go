package queue

type Queue []int //定义别名拓展类型

//v是一个可变参数的int类型
func (q *Queue) Push(v ...int) {
	//*q取值操作
	*q = append(*q, v...)
}

func (q *Queue) Pop() (tail int) {
	tail_index := len(*q) - 1
	tail = (*q)[tail_index]
	*q = (*q)[:tail_index]
	return
}

func (q *Queue) Shift() (head int) {
	head = (*q)[0]
	*q = (*q)[:0]

	return
}

func (q *Queue) Clear() {
	*q = (*q)[:0]
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
