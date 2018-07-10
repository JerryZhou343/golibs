// CopyRight (C) Jerry.Chau
// 非线安全,需要上层写日志保证,适配go-kit/kit 日志接口



package golibs

import (
	"os"
	"fmt"
	"path/filepath"
	"time"
	"strconv"
	"unsafe"
)

type logFile struct{
	curFile *os.File
	fileName string
	sizeFlag bool
	timeFlag bool
	filePath string
	sizeValue int64
	todayDate string
	msgQueue chan string
	closed bool
}

type  LogFileOption func(*logFile)
func NewLogFile(options ...LogFileOption) *logFile{

	logfile := logFile{
		fileName: "",
		sizeFlag: false,
		timeFlag: false,
		closed: false,
		msgQueue: make(chan string ,1000),
	}

	for _,option := range options{
		option(&logfile)
	}

	logfile.todayDate = time.Now().Format("2006-01-02")
	//
	if logfile.fileName != ""{
		file,err := os.OpenFile(logfile.filePath + logfile.fileName,
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil{
			fmt.Println(err.Error())
		}
		logfile.curFile = file
	}else {
		logfile.curFile = os.Stdout
	}

	go logfile.worker()

	return &logfile
}

//设置文件名
func LogFileName(fileName string) LogFileOption{
	return func(file *logFile){
		file.fileName = fileName
	}
}

//设置文件路径
func LogFilePath(path string)LogFileOption{
	return func(file *logFile){
		var slash string = string(os.PathSeparator)
		dir , _ := filepath.Abs(path)
		file.filePath = dir + slash
	}
}


//设置文件切割大小
func LogFileSize(size int, unit string)LogFileOption{
	return func(file *logFile){
		file.sizeFlag = true

		switch unit {
		case "M":
			file.sizeValue = int64(size) * 1024
		case "G":
			file.sizeValue = int64(size) * 1024 * 1024
		default:
			file.sizeValue = int64(size)
		}
	}
}

//按照天来切割
func LogFileTime(flag bool)LogFileOption{
	return func(file *logFile){
		file.timeFlag = true
	}
}


//
func (f *logFile)Write(p []byte) (n int, err error){
	str := (*string)(unsafe.Pointer(&p))
	f.msgQueue <- (*str)
	return len(p),nil
}

//切割文件
func (f *logFile)doRotate(){

	defer func(){
		rec := recover()
		if rec != nil{
			fmt.Println("doRotate %v", rec)
		}
	}()

	if f.curFile == nil{
		fmt.Println("doRotate curFile nil,return")
		return
	}
	//dir , _ := filepath.Abs(f.filePath)
	prefile := f.curFile
	_, err := prefile.Stat()

	if err == nil{
		filePath := f.filePath + f.fileName
		f.closed = true
		err := prefile.Close()
		if err != nil{
			fmt.Println("doRotate close err",err.Error())
		}
		nowTime := time.Now().Unix()
		err = os.Rename(filePath, filePath+"." + strconv.FormatInt(nowTime,10))
	}

	if f.fileName != ""{
		nextFile, err := os.OpenFile(f.filePath + f.fileName,
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil{
			fmt.Println(err.Error())
		}
		f.closed = false
		f.curFile = nextFile
		nowDate := time.Now().Format("2006-01-02")
		f.todayDate = nowDate
	}
}


func (f *logFile)worker(){
	for f.closed == false{
		msg := <- f.msgQueue
		fmt.Println(msg)
		f.curFile.WriteString(msg)
		if f.sizeFlag == true{
			curInfo,_ := os.Stat(f.filePath+f.fileName)
			if curInfo.Size() >= f.sizeValue {
				f.doRotate()
			}
		}
		nowDate := time.Now().Format("2006-01-02")
		if f.timeFlag == true &&
			nowDate != f.todayDate{
				f.doRotate()
		}
	}
}