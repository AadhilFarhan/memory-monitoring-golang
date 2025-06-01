package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("========== Memory Stats ==========")
	fmt.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\n", bToMb(m.Sys))
	fmt.Printf("Lookups = %v\n", m.Lookups)
	fmt.Printf("Mallocs = %v\n", m.Mallocs)
	fmt.Printf("Frees = %v\n", m.Frees)

	fmt.Println("------ Heap Stats ------")
	fmt.Printf("HeapAlloc = %v MiB\n", bToMb(m.HeapAlloc))
	fmt.Printf("HeapSys = %v MiB\n", bToMb(m.HeapSys))
	fmt.Printf("HeapIdle = %v MiB\n", bToMb(m.HeapIdle))
	fmt.Printf("HeapInuse = %v MiB\n", bToMb(m.HeapInuse))
	fmt.Printf("HeapReleased = %v MiB\n", bToMb(m.HeapReleased))
	fmt.Printf("HeapObjects = %v\n", m.HeapObjects)

	fmt.Println("------ GC Stats ------")
	fmt.Printf("NumGC = %v\n", m.NumGC)
	fmt.Println("==================================")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	
	go func() {
		fmt.Println("pprof server running at http://localhost:6060/debug/pprof/")
		http.ListenAndServe("localhost:6060", nil)
	}()

	for {
		printMemStats()
		time.Sleep(5 * time.Second)

		_ = make([]byte, 10*1024*1024) // Allocate 10MB
	}
}
