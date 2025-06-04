package workerpool

import "fmt"

// Worker обрабатывает входящие таски из пула
type Worker struct {
	id    int
	tasks <-chan string
	quit  <-chan struct{}
}

func (w *Worker) run() {
	for {
		select {
		case task, ok := <-w.tasks:
			if !ok {
				return
			}
			fmt.Printf("worker %d: %s\n", w.id, task)

		case <-w.quit:
			return
		}
	}
}
