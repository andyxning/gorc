package main

import (
	"fmt"
	"github.com/mr51m0n/gorc"
	"math/rand"
	"time"
)

var gorc0 gorc.Gorc

func main() {

	// no gorc here
	for i := 0; i < 20; i++ {
		go withoutgorc(i)
	}

	// with gorc this time
	for i := 0; i < 20; i++ {
		gorc0.Inc() // increase either before invoking a goroutine or within it
		go withgorc(i)
		gorc0.WaitLow(5) // no more then five goroutines governed by gorc0 are allowed at the same time
	}
}

func init() {
	gorc0.Init()
}

func withoutgorc(i int) {
	fmt.Println("Nr.", i, "goroutines without gorc")
	time.Sleep(time.Duration(rand.Int31n(2000)) * time.Millisecond)
}

func withgorc(i int) {
	defer gorc0.Dec() // decrease counter when finished
	fmt.Println("Nr.", i, " ", gorc0.Get(), "gorc goroutines running..")
	time.Sleep(time.Duration(rand.Int31n(2000)) * time.Millisecond)
}
