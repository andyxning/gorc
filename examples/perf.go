package main

import (
	"fmt"
	"github.com/mr51m0n/gorc"
	"math/rand"
	"time"
)

var gorc0 gorc.Gorc

func main() {
	// with gorc this time
	for i := 0; i < 100000; i++ {
		gorc0.Inc() // increase either before invoking a goroutine or within it
		go withgorc(i)
		gorc0.WaitLow(300) // no more than five goroutines governed by gorc0 are allowed at the same time
	}
}

func init() {
	gorc0.Init()
}

func withgorc(i int) {
	defer gorc0.Dec() // decrease counter when finished
	fmt.Println("Nr.", i, " ", gorc0.Get(), "gorc goroutines running..")
	time.Sleep(time.Duration(rand.Int31n(20)) * time.Millisecond)
}
