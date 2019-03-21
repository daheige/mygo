package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/daheige/thinkgo/common"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("fefefe") //默认输出到终端中

	//设置日志为json格式
	//log.SetFormatter(&log.JSONFormatter{})
	SetLogDir("./")
	JsonFormat(true)
	writeLog("info", map[string]interface{}{
		"age":     12,
		"message": "fefefe",
	})
}

const (
	PANIC = "panic" // 严重错误: 导致系统崩溃无法使用
	FATAL = "fatal" // 临界值错误: 超过临界值的错误，例如一天24小时，而输入的是25小时这样
	ERR   = "error" // 一般错误: 一般性错误
	WARN  = "warn"  // 警告性错误: 需要发出警告的错误
	INFO  = "info"  // 信息: 程序输出信息
	DEBUG = "debug" // 调试: 调试信息
	TRACE = "trace" //trace
)

var LogLevelMap = map[string]int{
	PANIC: 550,
	FATAL: 500,
	ERR:   400,
	WARN:  300,
	INFO:  200,
	DEBUG: 150,
	TRACE: 100,
}

var (
	logDir  = "" //日志文件存放目录
	logFile = "" //日志文件

	//采用sync实现加锁，效率比chan实现的加锁效率高一点
	//common.NewChanLock() 采用chan实现的乐观锁方式，实现加锁，效率稍微低一点
	logLock         = common.NewMutexLock()
	logTicker       = time.NewTicker(time.Second) //time一次性触发器
	logDay          = 0                           //当前日期
	logTimeZone     = "Local"                     //time zone default Local/Shanghai
	logtmFmtWithMS  = "2006-01-02 15:04:05.999"
	logtmFmtMissMS  = "2006-01-02 15:04:05"
	logtmFmtTime    = "2006-01-02"
	defaultLogLevel = INFO //默认为INFO级别
	logtmLoc        = &time.Location{}
)

//日志格式，是否以json格式保存
//默认非json格式
func JsonFormat(isJson bool) {
	if isJson {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

//设置日志记录时区
func SetLogTimeZone(timezone string) {
	logTimeZone = timezone
}

//日志存放目录
func SetLogDir(dir string) {
	if dir == "" {
		logDir = os.TempDir()
	} else {
		logDir = dir
	}

	logtmLoc, _ = time.LoadLocation(logTimeZone)
	now := time.Now().In(logtmLoc)

	//建立日志文件
	newFile(now)
}

//创建日志文件
func newFile(now time.Time) {
	if len(logDir) == 0 {
		return
	}

	logDay = now.Day()
	filename := filepath.Join(logDir, fmt.Sprintf("%s.%s.log", filepath.Base(os.Args[0]), now.Format(logtmFmtTime)))

	//创建文件
	fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(now.Format(logtmFmtMissMS), "open log file", filename, err, "use STDOUT")
	} else {
		fp.Close()
		logFile = filename
	}
}

func checkLogExist() {
	select {
	case <-logTicker.C:
	default:
		return
	}

	if !logLock.TryLock() {
		return
	}

	defer logLock.Unlock()

	loc, _ := time.LoadLocation(logTimeZone)
	now := time.Now().In(loc)
	//判断当天的日志文件是否存在，不存在就创建
	if logDay != now.Day() {
		newFile(now)
	}
}

func writeLog(levelName string, fields log.Fields) {
	if len(fields) < 1 {
		return
	}

	if _, ok := LogLevelMap[levelName]; !ok {
		levelName = defaultLogLevel
	}

	var msg interface{}
	if _, ok := fields["message"]; ok {
		msg = fields["message"]
		delete(fields, "message")
	}

	//fmt.Println("msg: ", msg)

	//检测日志文件是否存在
	checkLogExist()

	//函数调用位置
	_, file, line, _ := runtime.Caller(2)
	now := time.Now().In(logtmLoc)

	fields["trace_file"] = filepath.Base(file)
	fields["trace_line"] = line
	fields["trace_time"] = now.Format(logtmFmtWithMS)

	//开始写日志，这里需要对文件句柄进行加锁
	logLock.Lock()
	defer logLock.Unlock()

	fp, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		now := time.Now().In(logtmLoc)
		fmt.Println(now.Format(logtmFmtMissMS), "open log file", logFile, err)
		fmt.Println("fields: ",fields)
		return
	}

	defer fp.Close()

	log.SetOutput(fp)

	entry := log.WithFields(fields)

	switch levelName {
	case "info":
		entry.Infoln(msg)
	case "debug":
		entry.Debugln(msg)
	case "error":
		entry.Errorln(msg)
	case "warn":
		entry.Warnln(msg)
	case "trace":
		entry.Traceln(msg)
	case "fatal":
		entry.Fatalln(msg)
	case "panic":
		entry.Panicln(msg)
	default:
		entry.Infoln(msg)
	}
}

func Info(m map[string]interface{}){
	writeLog("info",m)
}

func Debug(m map[string]interface{}){
	writeLog("debug",m)
}

func Error(m map[string]interface{}){
	writeLog("error",m)
}

func Trace(m map[string]interface{}){
	writeLog("trace",m)
}

func Fatal(m map[string]interface{}){
	writeLog("fatal",m)
}

func Panic(m map[string]interface{}){
	writeLog("panic",m)
}

