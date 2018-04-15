package entry

//同一个目录下的文件,包名必须相同
//工厂函数返回结构体的地址
func CreateNode(value int) *TreeNode {
	//返回局部变量的地址
	return &TreeNode{Value: value}
}
