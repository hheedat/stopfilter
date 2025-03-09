package util

import (
	"runtime"
	"github.com/hheedat/stopfilter/log"
	"strconv"
)

func PrintMem() {
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	log.Log.Info("mem.sys " + strconv.FormatUint(memStats.Sys, 10))
}
