package main

import (
	"fmt"
	"goweb/slog"
	"strconv"
	"time"
	// "os"
	// "path/filepath"
)

const (
	kb = 1 << (10 * iota)
	mb
	gb
	tb
	pb
)

func main() {
	// var t time.Time
	s := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(s)
	fmt.Println(time.Now().Format("2006_01_02_15_04"))
	// res := filepath.Dir("/runtime/abc/index.php") //获取目录名
	// str := filepath.Base("/etc/hosts/ab") //获取文件名
	// fmt.Println(str)
	// fmt.Println(res)

	// file, err := os.Stat("hg_test.log")
	// fmt.Println(err)
	// fmt.Println(file.Size())
	// slog.Debug("ssss")
	slog.InitLogSize(1024)
	slog.Info("tess232", "test/hg_test")
	// time.Sleep(2 * time.Second)

	fmt.Println(kb, mb, gb, tb, pb)
	// slog.Error("sss123", "hg_error")
	fmt.Println("fefefe")
}
