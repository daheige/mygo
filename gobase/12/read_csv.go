package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type DataList struct {
	uid   uint64
	price float64
	desc  string
}

func main() {
	datas := make([]DataList, 1)
	fp, err := os.Open("./demo.csv") //打开文件
	if err != nil {
		log.Fatalf("Error %s open the csv", err)
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)

	for {
		line, err := reader.ReadString('\n') //读取字符串到一行中
		//读取数据结束
		if err == io.EOF {
			break
		}

		line = string(line[:len(line)-2]) //去掉\r\n
		strData := strings.Split(line, ",")
		if strData[0] == "uid" { //第一行表头不需要
			continue
		}

		// fmt.Println(line)
		data := new(DataList)
		data.uid, _ = strconv.ParseUint(strData[0], 10, 64) //转换为64位无符号整数
		data.price, _ = strconv.ParseFloat(strData[1], 64)
		data.desc = strData[2]

		if datas[0].uid == 0 {
			datas[0] = *data
		} else {
			datas = append(datas, *data)
		}
	}
	fmt.Println(datas)

	for _, v := range datas {
		fmt.Println(v.uid, v.price, v.desc)
	}

}
