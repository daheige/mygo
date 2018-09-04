package main

import "fmt"

func main()  {
	/**
	ch初始化后，case1读取失败，timeout同样失败，因为channel中无数据，直接跳至default执行并返回。
	注意如果没有default，select 会一直等待等到某个 case 语句完成
	也就是等到成功从 ch 或者 timeout 中读到数据
	否则一直阻塞。
	 */
	t := make(chan bool,1)
	ch := make(chan int)
	select{
	case <-ch:
		case <-t:
			fmt.Println("timeout")
		default:
			fmt.Println("default case is running")
	}
}
