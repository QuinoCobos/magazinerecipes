package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	u "magazinerecipes/utils"
)

//Logger variables
var (
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	sqlW    *log.Logger
	errW    *log.Logger
	level   int
)

//Init Initializes loggers
func Init(logLevel int) {
	level = logLevel
	debugHandle, infoHandle, warningHandle, sqlHandle, errorHandle := setWriters()

	switch level {
	case 3:
		debug = log.New(debugHandle,
			"DEBUG: ",
			log.Ldate|log.Ltime)
		fallthrough
	case 2:
		info = log.New(infoHandle,
			"INFO: ",
			log.Ldate|log.Ltime)
		fallthrough
	case 1:
		warning = log.New(warningHandle,
			"WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile)
		fallthrough
	default:
		sqlW = log.New(sqlHandle,
			"SQL: ",
			log.Ldate|log.Ltime|log.Lshortfile)
		errW = log.New(errorHandle,
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func Debug(str string) {
	if level == 3 {
		debug.Println(str)
	}
}

func Info(str string) {
	if level >= 2 {
		info.Println(str)
	} else {
		fmt.Println(str)
	}
}

func Warning(str string) {
	if level >= 1 {
		warning.Println(str)
	} else {
		fmt.Println(str)
	}
}

func Sql(str string) {
	sqlW.Println(str)
}

func Error(str string) {
	errW.Println(str)
}

func setWriters() (io.Writer, io.Writer, io.Writer, io.Writer, io.Writer) {
	dir, err := os.UserHomeDir()
	u.Check(err)
	os.MkdirAll(dir+u.OutputPath, 2777)
	infoFile, infoErr := os.OpenFile(dir+u.OutputPath+time.Now().Format(u.DateFormat)+"-recipes.log", os.O_CREATE|os.O_APPEND, 2777)
	u.Check(infoErr)
	sqlFile, sqlErr := os.OpenFile(dir+u.OutputPath+time.Now().Format(u.DateFormat)+"-sql.log", os.O_CREATE|os.O_APPEND, 2777)
	u.Check(sqlErr)
	errFile, errErr := os.OpenFile(dir+u.OutputPath+time.Now().Format(u.DateFormat)+"-error.log", os.O_CREATE|os.O_APPEND, 2777)
	u.Check(errErr)
	infoLog := io.MultiWriter(infoFile, os.Stdout)
	errorLog := io.MultiWriter(errFile, os.Stderr)
	return infoFile, infoLog, infoLog, sqlFile, errorLog
}
