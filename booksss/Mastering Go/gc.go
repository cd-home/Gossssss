package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	f, err := os.Create("./gcTrace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		log.Println(err)
		return
	}
	defer trace.Stop()

	var mem runtime.MemStats
	PrintGcToFile(mem)

	for i := 0; i < 10; i++ {
		_ = make([]byte, 500000000)
	}
	PrintGcToFile(mem)
}

func PrintGcToFile(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	log.Println("mem Alloc", mem.Alloc)
	log.Println("mem TotalAlloc", mem.TotalAlloc)
	log.Println("mem HeapAlloc", mem.HeapAlloc)
	log.Println("mem NumGc", mem.NumGC)
}