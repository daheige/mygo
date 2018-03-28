package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("fefe")
	if content, err := ioutil.ReadFile("a.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content))
	}

	file, err := os.Open("a.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	//浏览其中的内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //相当于while
		fmt.Println(scanner.Text())
	}

	//无限循环不需要条件
	for {
		fmt.Println("abc")
		break
	}

}
