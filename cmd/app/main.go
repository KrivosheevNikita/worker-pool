package main

import (
	"fmt"
	"log"

	"github.com/KrivosheevNikita/worker-pool/internal/workerpool"
)

func main() {
	pool := workerpool.New(3)

	for i := 0; i < 10; i++ {
		pool.AddTask(fmt.Sprintf("task %d", i))
	}

	pool.AddWorker()
	pool.AddWorker()
	log.Printf("workers count: %d", pool.CountWorkers())

	pool.DeleteWorker()
	log.Printf("workers count: %d", pool.CountWorkers())

	pool.Stop()
}
