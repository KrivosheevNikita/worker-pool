package workerpool

import "sync"

// Pool управляет набором воркеров и очередью задач
type Pool struct {
	tasks   chan string
	workers map[int]chan struct{}
	closed  bool
	lastID  int

	wg sync.WaitGroup
	mu sync.Mutex
}

// New создаёт пул из n воркеров
func New(n int) *Pool {
	p := &Pool{
		tasks:   make(chan string),
		workers: make(map[int]chan struct{}),
	}
	for i := 0; i < n; i++ {
		p.AddWorker()
	}
	return p
}

// AddWorker добавляет воркера в пул
func (p *Pool) AddWorker() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return
	}

	p.lastID++
	id := p.lastID
	quit := make(chan struct{})
	p.workers[id] = quit

	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		w := &Worker{id: id, tasks: p.tasks, quit: quit}
		w.run()
	}()
}

// DeleteWorker удаляет одного воркера из пула
func (p *Pool) DeleteWorker() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for id, quit := range p.workers {
		close(quit)
		delete(p.workers, id)
		return
	}
}

// AddTask ставит строку в очередь заданий
func (p *Pool) AddTask(task string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return
	}
	p.tasks <- task
}

// Stop завершает работу пула
func (p *Pool) Stop() {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return
	}
	p.closed = true
	close(p.tasks)
	for _, quit := range p.workers {
		close(quit)
	}
	p.mu.Unlock()

	p.wg.Wait()
}

// CountWorkers возвращает текущее количество воркеров
func (p *Pool) CountWorkers() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.workers)
}
