package main

import (
	"fmt"
	"runtime"
)

func main() {
	var memStats runtime.MemStats
	runtime.ReadMemStats((&memStats))
	fmt.Printf("Free Mem: %v; Sys Mem: %v\n", memStats.Frees, memStats.Sys)
	fmt.Printf("Number of Cpus: %v; Number of threads: %v\n", runtime.NumCPU(), runtime.GOMAXPROCS(0))
}
