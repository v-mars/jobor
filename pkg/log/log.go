package log

import (
	"io"
	"log"
	"os"
)


var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

type Logger struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func Init(){
	//errFile,err:=os.OpenFile("app.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	logPath := "./logs"

	logFile,err:=os.OpenFile(logPath+"/app.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	errFile,err:=os.OpenFile(logPath+"/app_error.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	if err!=nil{
		log.Fatalln("打开日志文件失败：",err)
	}

	Info = log.New(io.MultiWriter(os.Stdout, logFile),"[INFO] ",log.Ldate | log.Ltime | log.Lshortfile)
	Warning = log.New(io.MultiWriter(os.Stdout, logFile),"[WARNING] ",log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr,logFile, errFile),"[ERROR] ",log.Ldate | log.Ltime | log.Lshortfile)

}

//Info.Println("ocean:","http://www.flysnow.org")
//Warning.Printf("ocean：%s\n","flysnow_org")
//Error.Println("欢迎关注留言")
