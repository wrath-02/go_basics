// Topic 09: Channels and select.
//
// Run: go run ./09-channels
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	unbuffered()
	buffered()
	closeAndRange()
	directional()
	selectStatement()
	timeouts()
	workerPool()
	dealFinder()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

func unbuffered() {
	section("Unbuffered channel: send blocks until someone receives")
	ch := make(chan int) // a channel that carries ints

	// Sending on an unbuffered channel blocks until another goroutine receives.
	// Doing "ch <- 1" and "<-ch" in the SAME goroutine would deadlock:
	//   fatal error: all goroutines are asleep - deadlock!
	go func() {
		ch <- 42 // send
	}()
	value := <-ch // receive (blocks until the value arrives)
	fmt.Println("received:", value)
}

func buffered() {
	section("Buffered channel: send only blocks when the buffer is full")
	ch := make(chan string, 2)
	ch <- "first" // doesn't block, buffer has room
	ch <- "second"
	fmt.Println("len:", len(ch), "cap:", cap(ch))
	fmt.Println(<-ch, <-ch) // FIFO order
}

func produce(n int, ch chan int) {
	defer close(ch) // close tells receivers "no more values are coming"
	for i := range n {
		ch <- i * i
	}
}

func closeAndRange() {
	section("close() and range over a channel")
	ch := make(chan int)
	go produce(5, ch)
	for v := range ch { // loop ends automatically when ch is closed
		fmt.Print(v, " ")
	}
	fmt.Println()

	// Receiving from a closed channel returns the zero value immediately.
	v, ok := <-ch
	fmt.Println("after close:", v, "ok =", ok)
}

// chan<- int is send-only and <-chan int is receive-only.
// The compiler then stops you from using a channel the wrong way.
func sender(out chan<- int) {
	for i := 1; i <= 3; i++ {
		out <- i
	}
	close(out)
}

func receiver(in <-chan int, done chan<- int) {
	total := 0
	for v := range in {
		total += v
	}
	done <- total
}

func directional() {
	section("Directional channels")
	nums := make(chan int)
	done := make(chan int)
	go sender(nums)
	go receiver(nums, done)
	fmt.Println("sum from receiver:", <-done)
}

func selectStatement() {
	section("select: wait on several channels at once")
	fast := make(chan string)
	slow := make(chan string)
	go func() { time.Sleep(10 * time.Millisecond); fast <- "fast" }()
	go func() { time.Sleep(50 * time.Millisecond); slow <- "slow" }()

	for range 2 {
		select { // runs whichever case is ready first
		case msg := <-fast:
			fmt.Println("received", msg)
		case msg := <-slow:
			fmt.Println("received", msg)
		}
	}

	// A default case makes select non-blocking.
	empty := make(chan int)
	select {
	case v := <-empty:
		fmt.Println("got", v)
	default:
		fmt.Println("nothing ready, moving on (default case)")
	}
}

func timeouts() {
	section("Timeouts with select + time.After")
	result := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond) // too slow
		result <- "done"
	}()

	select {
	case r := <-result:
		fmt.Println(r)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("timed out after 50ms")
	}
}

func workerPool() {
	section("Worker pool: N workers share one jobs channel")
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup
	for range 3 { // 3 workers
		wg.Go(func() {
			for j := range jobs {
				results <- j * 10
			}
		})
	}

	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs) // workers exit their range loop once jobs is drained

	wg.Wait()
	close(results)

	sum := 0
	for r := range results {
		sum += r
	}
	fmt.Println("sum of results:", sum, "(10+20+...+60 = 210)")
}

// dealFinder watches several shops' prices and reacts to the first deal.
const maxLaptopPrice = 800.0

func checkShop(shop string, prices []float64, deals chan<- string) {
	for _, p := range prices {
		time.Sleep(10 * time.Millisecond) // pretend to fetch the price
		if p <= maxLaptopPrice {
			deals <- fmt.Sprintf("%s has a laptop for $%.2f", shop, p)
			return
		}
	}
}

func dealFinder() {
	section("Example: first deal wins")
	deals := make(chan string)
	go checkShop("ShopA", []float64{999, 950, 920, 899, 870}, deals) // never a deal
	go checkShop("ShopB", []float64{1100, 980, 790}, deals)          // deal on 3rd check
	go checkShop("ShopC", []float64{900, 880, 860, 840, 799}, deals) // deal on 5th check

	select {
	case msg := <-deals:
		fmt.Println("Deal found:", msg)
	case <-time.After(time.Second):
		fmt.Println("No deals today")
	}
}
