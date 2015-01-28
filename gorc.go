package gorc

import (
	"sync"
	"time"
)

type Gorc struct {
	count      int
	waitMillis time.Duration
	sync.Mutex
}

// Inc increases the counter by one.
func (g *Gorc) Inc() {
	g.Lock()
	g.count++
	g.Unlock()
}

// IncBy increases the counter by b.
func (g *Gorc) IncBy(b int) {
	g.Lock()
	g.count += b
	g.Unlock()
}

// Dec decreases the counter by one.
func (g *Gorc) Dec() {
	g.Lock()
	g.count--
	g.Unlock()
}

// DecBy decreases the counter by b.
func (g *Gorc) DecBy(b int) {
	g.Lock()
	g.count -= b
	g.Unlock()
}

// GetCount returns an integer holding the count.
func (g *Gorc) Get() int {
	return int(g.count)
}

// SetWaitMillis sets the time in milliseconds the Wait function
// waits between checking the count against the given integer.
func (g *Gorc) SetWaitMillis(w int) {
	g.Lock()
	g.waitMillis = time.Duration(w) * time.Millisecond
	g.Unlock()
}

// Init initializes a new Gorc instance
func (g *Gorc) Init() {
	g.Lock()
	g.count = 0
	g.waitMillis = 100 * time.Millisecond
	g.Unlock()
}

// WaitLow will return as soon as the Gorc counter falls below w.
// e.g. wait until all but w goroutines are stopped.
func (g *Gorc) WaitLow(w int) {
	for g.count >= w {
		time.Sleep(g.waitMillis)
	}
	return
}

// WaitHigh will return as soon as the Gorc counter goes above w.
// e.g. wait until at least w goroutines are started.
func (g *Gorc) WaitHigh(w int) {
	for g.count <= w {
		time.Sleep(g.waitMillis)
	}
	return
}
