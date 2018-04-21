package main

import (
	"log"
	"os"
	"fmt"
	"time"
)

const (
	LevelError = iota
	LevelWarn
	LevelInfo
	LevelDebug
)

type Logger struct {
	l         *log.Logger
	level     int
	callDepth int
}

var Log Logger
var f *os.File

var date = ""
var dateFormat = "2006-01-02"

//func main() {
//
//	i := 1
//	for {
//		Log.Error("000")
//		Log.Warn("111")
//		Log.Info("222")
//		Log.Debug("333")
//		i++
//		if i > 10000000 {
//			break
//		}
//		time.Sleep(500 * time.Millisecond)
//	}
//
//	Destroy()
//}

func init() {
	initLogger()
}

func initLogger() {
	checkRotate()

	Log.level = 2
	Log.callDepth = 2
}

func checkRotate() {
	nowDate := getFormatDate()
	if nowDate != date {

		if f != nil {
			rotateLog()
		}

		path := getLogPath()
		date = nowDate

		openFile(path)

		Log.l = log.New(f, "", log.Lshortfile|log.Ldate|log.Lmicroseconds)
	}
}

func Destroy() {
	closeFile()
}

func openFile(path string) {
	var err error

	f, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func closeFile() {
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func getLogPath() string {
	path := Conf.LogPath
	return path
}

func rotateLog() {
	closeFile()
	oldName := getLogPath()
	newName := oldName + time.Now().Add(-time.Second * 86400).Format(dateFormat)
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

func getFormatDate() string {
	return time.Now().Format(dateFormat)
}

func (ll Logger) Error(v interface{}) {
	if ll.level < LevelError {
		return
	}
	checkRotate()

	ll.l.SetPrefix("[Error] : ")
	ll.l.Output(ll.callDepth, fmt.Sprintln(v))
}

func (ll Logger) Warn(v interface{}) {
	if ll.level < LevelWarn {
		return
	}
	checkRotate()

	ll.l.SetPrefix("[Warn] : ")
	ll.l.Output(ll.callDepth, fmt.Sprintln(v))
}

func (ll Logger) Info(v interface{}) {
	if ll.level < LevelInfo {
		return
	}
	checkRotate()

	ll.l.SetPrefix("[Info] : ")
	ll.l.Output(ll.callDepth, fmt.Sprintln(v))
}

func (ll Logger) Debug(v interface{}) {
	if ll.level < LevelDebug {
		return
	}
	checkRotate()

	ll.l.SetPrefix("[Debug] : ")
	ll.l.Output(ll.callDepth, fmt.Sprintln(v))
}
