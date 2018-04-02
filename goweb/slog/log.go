package slog

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	kb = 1 << (10 * iota) //字节大小
	mb
	gb
	tb
	pb
)

var (
	logLevel     = log.LstdFlags                           //标准日志格式
	logWriteType = os.O_CREATE | os.O_WRONLY | os.O_APPEND //文件创建标识
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func Save(message interface{}, filename string, prefix string) {
	var isExist = true
	filename = filename + ".log"
	path := filepath.Dir(filename) //得到日志目录

	if file, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			isExist = false
		}
	} else {
		//切割日志,备份操作
		if file.Size() > 2*mb {
			newfile := strconv.FormatInt(time.Now().Unix(), 10) + ".log"
			os.Rename(filename, path+"/"+newfile)
			os.Create(filename)
		}
	}

	filename = filepath.Base(filename) //得到日志文件名
	if !isExist {
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Fatalf("创建目录%s失败,错误信息: %s", path, err)
		}
	}

	//打开日志文件
	errFile, err := os.OpenFile(filename, logWriteType, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	var Info *log.Logger
	Info = log.New(io.MultiWriter(os.Stderr, errFile), prefix, logLevel) //第二个参数是日志内容前缀
	Info.Println(message)
}

func Debug(message interface{}) {
	var debug *log.Logger
	debug = log.New(os.Stdout, "Debug: ", logLevel)
	debug.Println(message)
}

func Warn(message interface{}, filename string) {
	Save(message, filename, "warn: ")
}

func Info(message interface{}, filename string) {
	Save(message, filename, "Info: ")
}

func Error(message interface{}, filename string) {
	Save(message, filename, "Error: ")
}
