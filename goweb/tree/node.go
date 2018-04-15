package main

import (
	"fmt"
	"goweb/tree/entry"
)

//拓展已有的entry.TreeNode
type myTreeNode struct {
	node *entry.TreeNode
}

//从后面遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{
		myNode.node.Left,
	}

	left.postOrder()

	right := myTreeNode{
		myNode.node.Right,
	}

	right.postOrder()

	myNode.node.Print()

}

func main() {
	var root entry.TreeNode

	root = entry.TreeNode{Value: 3}
	root.Left = &entry.TreeNode{}
	root.Left.Right = entry.CreateNode(2)
	root.Right = &entry.TreeNode{
		5,
		nil,
		nil,
	}
	root.Right.Left = new(entry.TreeNode) //重新赋值

	fmt.Println(root.Left.Right)

	nodes := []entry.TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	root.Print()

	root.Right.Left.SetValue(4)
	root.Right.Left.Print() //并没有改变值
	root.Right.Left.ChangeValue(4)
	root.Right.Left.Print() //4

	var pRoot *entry.TreeNode
	pRoot.ChangeValue(200)
	pRoot = &root

	pRoot.ChangeValue(3000) //root也跟着改变了,因为pRoot和root指向同一个内存地址
	pRoot.Print()
	root.Print()
	fmt.Println("=====开始遍历root====")
	root.List() //0 2 3000 4 5
	fmt.Println("-----")

	myRoot := myTreeNode{&root}
	myRoot.postOrder()

}
