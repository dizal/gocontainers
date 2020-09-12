package pool

import (
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	size := 10
	pool := New(size)

	var count int
	var m sync.Mutex

	for i := 0; i < 20; i++ {
		pool.Add()

		m.Lock()
		count++
		m.Unlock()

		if count > size {
			t.Fatalf("Max goroutines > %v [%v]", size, count)
		}

		go func() {
			time.Sleep(10 * time.Millisecond)

			m.Lock()
			count--
			m.Unlock()

			pool.Done()
		}()
	}

	pool.Wait()
}
