# gorc
set thresholds on number of running goroutines

Can increase and decrease a counter when starting or stopping a goroutine. It can wait for a minimum or maximum number of goroutines running, thus allowing to set thresholds for the number of gorc governed goroutines running at the same time.

minimalist docs at https://godoc.org/github.com/mr51m0n/gorc


`import "github.com/mr51m0n/gorc"`

## an example:

```Go
var gorc0 gorc.Gorc

func main() {
	for i := 0; i < 20; i++ {
		gorc0.Inc()
		go withgorc(i)
		gorc0.WaitLow(5) // no more than five goroutines governed by gorc0 are allowed at the same time
	}
}

func init() {
	gorc0.Init()
}

func withgorc(i int) {
	defer gorc0.Dec() // decrease counter when finished
	fmt.Println("Nr.", i, " ", gorc0.Get(), "gorc goroutines running..")
	time.Sleep(time.Duration(rand.Int31n(2000)) * time.Millisecond)
}
```
