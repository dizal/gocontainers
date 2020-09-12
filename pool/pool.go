package pool

import "sync"

// Pool implements a simple goruntine pool
type Pool struct {
	wg   sync.WaitGroup
	pool chan byte
}

// New creates a waitgroup with a specific size (the maximum number of
// goroutines to run at the same time).
func New(size int) Pool {
	return Pool{
		pool: make(chan byte, size),
	}
}

// Add pushes ‘one’ into the group. Blocks if the group is full
func (p *Pool) Add() {
	p.pool <- 1
	p.wg.Add(1)
}

// Done pops ‘one’ out the group
func (p *Pool) Done() {
	<-p.pool
	p.wg.Done()
}

// Wait waiting the group empty
func (p *Pool) Wait() {
	p.wg.Wait()
}
