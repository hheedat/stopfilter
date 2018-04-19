package main

import (
	"log"
	"os"
	"fmt"
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

func main() {
	Log.Error("000")
	Log.Warn("111")
	Log.Info("222")
	Log.Debug("333")
	Close()
}

func init() {
	path := "/var/logs/error.log"

	var err error

	f, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}

	Log.l = log.New(f, "", log.Lshortfile|log.Lmicroseconds)

	Log.level = 2
	Log.callDepth = 2
}

func (ll Logger) Error(v interface{}) {
	if ll.level < LevelError {
		return
	}

	ll.l.SetPrefix("[Error] : ")
	ll.l.Output(ll.callDepth, fmt.Sprintln(v))
}

func (ll Logger) Warn(v interface{}) {
	if ll.level < LevelWarn {
		return
	}

	ll.l.SetPrefix("[Warn] : ")
	ll.l.Output(ll.callDepth, fmt.Sprintln(v))
}

func (ll Logger) Info(v interface{}) {
	if ll.level < LevelInfo {
		return
	}

	ll.l.SetPrefix("[Info] : ")
	ll.l.Output(ll.callDepth, fmt.Sprintln(v))
}

func (ll Logger) Debug(v interface{}) {
	if ll.level < LevelDebug {
		return
	}

	ll.l.SetPrefix("[Debug] : ")
	ll.l.Output(ll.callDepth, fmt.Sprintln(v))
}

func Close() {
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
