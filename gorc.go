package gorc

import (
	"sync/atomic"
	"time"
	"errors"
)

var (
	ErrDecreasedByNegative = errors.New("Decreased by a negative number")
	ErrIncreasedByNegative = errors.New("Increased by a negative number")
)

type Gorc struct {
	count      int32
	waitMillis int64
}

// Inc increases the counter by one.
func (g *Gorc) Inc() {
	atomic.AddInt32(&g.count, 1)
}

// IncBy increases the counter by b.
// b must be a positive number, otherwise an ErrIncreasedByNegative will be
// returned.
func (g *Gorc) IncBy(b int32) error {
	if b < 0 {
		return ErrIncreasedByNegative
	}
	atomic.AddInt32(&g.count, b)
	return nil
}

// Dec decreases the counter by one.
func (g *Gorc) Dec() {
	atomic.AddInt32(&g.count, -1)
}

// DecBy decreases the counter by b.
// b must be a positive number, otherwise an ErrDecreasedByNegative will be
// returned.
func (g *Gorc) DecBy(b int32) error {
	if b < 0 {
		return ErrDecreasedByNegative
	} else {
		b = int32(^uint32(b-1))
	}
	atomic.AddInt32(&g.count, b)
	return nil
}

// GetCount returns an integer holding the count.
func (g *Gorc) Get() int32 {
	return atomic.LoadInt32(&g.count)
}

// SetWaitMillis sets the time in milliseconds the Wait function
// waits between checking the count against the given integer.
func (g *Gorc) SetWaitMillis(w int64) {
	atomic.StoreInt64(&g.waitMillis, w)
}

// Init initializes a new Gorc instance
func (g *Gorc) Init() {
	atomic.StoreInt32(&g.count, 0)
	atomic.StoreInt64(&g.waitMillis, 100)
}

// WaitLow will return as soon as the Gorc counter falls below w.
// e.g. wait until all but w goroutines are stopped.
func (g *Gorc) WaitLow(w int32) {
	for atomic.LoadInt32(&g.count) >= w {
		dur := time.Duration(atomic.LoadInt64(&g.waitMillis))
		time.Sleep(dur * time.Millisecond)
	}
	return
}

// WaitHigh will return as soon as the Gorc counter goes above w.
// e.g. wait until at least w goroutines are started.
func (g *Gorc) WaitHigh(w int32) {
	for atomic.LoadInt32(&g.count) <= w {
		dur := time.Duration(atomic.LoadInt64(&g.waitMillis))
		time.Sleep(dur * time.Millisecond)
	}
	return
}
