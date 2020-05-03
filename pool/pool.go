package util

import (
	"sync"
	"time"
)

// DefaultSleepTime ...
const DefaultSleepTime = 50 * time.Millisecond

// Pool ...
type Pool struct {
	sync.Mutex
	max   uint8
	count uint8
	sleep time.Duration
}

// New ...
func New(max uint8, sleep time.Duration) Pool {
	return Pool{
		max:   max,
		sleep: sleep,
	}
}

// Add ...
func (p *Pool) Add() {
	p.Lock()
	p.count++
	p.Unlock()
}

// Done ...
func (p *Pool) Done() {
	p.Lock()
	p.count--
	p.Unlock()
}

// Wait ...
func (p *Pool) Wait() {
	for {
		if p.count <= p.max {
			break
		} else {
			time.Sleep(p.sleep)
		}
	}
}

// WaitEmpty ...
func (p *Pool) WaitEmpty() {
	for {
		if p.count > 0 {
			time.Sleep(p.sleep)
		} else {
			break
		}
	}
}
