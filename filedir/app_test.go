package filedir

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

/*func Test_mkdir(t *testing.T) {
	//创建文件夹
	os.Mkdir("heige", 0777)
	os.MkdirAll("heige/test1/test2", 0755) //递归创建目录
	err := os.Remove("heige/test1/test2")  //删除名称为name的目录，当目录下有文件或者其他目录是会出错
	if err != nil {
		fmt.Println(err)
	}

	os.RemoveAll("heige") //删除所有目录 path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除
	fmt.Println("ok")
	t.Log("ok")
}*/

//创建文件
func Test_createFile(t *testing.T) {
	filename := "heige.md"
	fout, err := os.Create(filename)
	if err != nil {
		fmt.Println(filename, err)
		return
	}

	defer fout.Close() //退出函数后关闭文件操作句柄

	for i := 0; i < 10; i++ {
		fout.WriteString("just a test\r\n")
		fout.Write([]byte("just a test" + strconv.Itoa(i) + "\r\n"))
	}

}

func Test_readFile(t *testing.T) {
	filename := "heige.txt"
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Println(filename, err)
		return
	}

	defer fd.Close()
	fmt.Println(strconv.Itoa(12))
	buf := make([]byte, 1024) //创建字节切片，容量是1024
	for {
		n, _ := fd.Read(buf) //读取数据到buf中

		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func Test_delFile(t *testing.T) {
	os.Rename("heige.txt", "test_heige.md") //重命名
	os.Remove("heige.md")
	// os.RemoveAll("heige.md")
}
