package concurrent

import (
	"io"
	"sync"
)

type ChanWriter struct {
	stdout chan []byte
}

func (c *ChanWriter) Write(p []byte) (int, error) {
	c.stdout <- append([]byte{}, p...) // Send a copy
	return len(p), nil
}

func (c *ChanWriter) Close() error {
	if c.stdout != nil {
		close(c.stdout)
	}
	return nil
}

type ParallelLogGroup struct {
	queue       chan chan []byte
	stdoutMutex sync.Mutex
	stdout      io.Writer
}

func NewParallelLogGroup(w io.Writer) *ParallelLogGroup {
	return &ParallelLogGroup{
		queue:       make(chan chan []byte, 8),
		stdoutMutex: sync.Mutex{},
		stdout:      w,
	}
}

func (p *ParallelLogGroup) Logger() ChanWriter {
	c := make(chan []byte, 2048)
	p.queue <- c
	cw := ChanWriter{c}
	return cw
}

func (p *ParallelLogGroup) StartOutputStream() {
	go p.stream()
}

func (p *ParallelLogGroup) StopOutputStream() {
	if p.queue != nil {
		close(p.queue)
	}

	p.queue = make(chan chan []byte, 8)
	p.stdoutMutex = sync.Mutex{}
}

func (p *ParallelLogGroup) stream() {
	for log := range p.queue {
		for word := range log {
			p.stdoutMutex.Lock()
			p.stdout.Write(word)
			p.stdoutMutex.Unlock()
		}
	}
}
