package main

import (
	"runtime"
	"fmt"
)

func PrintMem() {
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	fmt.Println("mem.sys ", memStats.Sys)
}
