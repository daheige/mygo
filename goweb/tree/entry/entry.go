package entry

import (
	"fmt"
)

//包导出的函数名和结构体字段必须是大写
//定义结构体
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

//为结构体定义方法
//func后面是接受者,然后是函数名称
func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

//如果需要改变Value,方法接收者应该是一个指针
func (node TreeNode) SetValue(value int) {
	node.Value = value
}

func (node *TreeNode) ChangeValue(value int) {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	node.Value = value
}

func (node *TreeNode) List() {
	if node == nil {
		return
	}

	node.Left.List()
	node.Print()
	node.Right.List()
}
