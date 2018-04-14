package main

import (
	"fmt"
)

//定义结构体
type treeNode struct {
	value       int
	left, right *treeNode
}

//为结构体定义方法
//func后面是接受者,然后是函数名称
func (node treeNode) print() {
	fmt.Println(node.value)
}

//如果需要改变value,方法接收者应该是一个指针
func (node treeNode) setValue(value int) {
	node.value = value
}

func (node *treeNode) changeVaule(value int) {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	node.value = value
}

func (node *treeNode) list() {
	if node == nil {
		return
	}

	node.left.list()
	node.print()
	node.right.list()

}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.left.right = createNode(2)
	root.right = &treeNode{
		5,
		nil,
		nil,
	}
	root.right.left = new(treeNode) //重新赋值

	fmt.Println(root.left.right)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	root.print()

	root.right.left.setValue(4)
	root.right.left.print() //并没有改变值
	root.right.left.changeVaule(4)
	root.right.left.print() //4

	var pRoot *treeNode
	pRoot.changeVaule(200)
	pRoot = &root

	pRoot.changeVaule(3000) //root也跟着改变了,因为pRoot和root指向同一个内存地址
	pRoot.print()
	root.print()
	fmt.Println("=====开始遍历root====")
	root.list() //0 2 3000 4 5

}

//工厂函数返回结构体的地址
func createNode(value int) *treeNode {
	//返回局部变量的地址
	return &treeNode{value: value}
}
