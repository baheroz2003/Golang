// In Go (Golang), channels are a powerful feature used for communication between goroutines. They provide a way to safely share data without explicit locks or condition variables.
package main // yeh main package declare kar raha hai, jo Go program ka entry point hota hai
import (
	"fmt"  // fmt package input/output ke liye use hota hai (jaise print)
	"sync" // sync package concurrency tools jaise WaitGroup ke liye use hota hai
)
func main() {
	fmt.Println("channels in golang") // Console par message print karega

	myCh := make(chan int, 2) // ek buffered channel banaya jiska type int hai aur capacity 2 hai

	wg := &sync.WaitGroup{} // ek WaitGroup banaya taaki hum goroutines ke complete hone ka wait kar sakein
	wg.Add(2)               // 2 goroutines ka wait karna hai, isliye 2 add kiya

	// 👇 Ye goroutine sirf receive karega channel se
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// val, ischanelopen mein value aur channel ka status store hoga
		val, ischanelopen := <-myCh // channel se ek value receive karo aur check karo channel open hai ya nahi
		fmt.Println(ischanelopen)   // yeh print karega true ya false (channel open hai ya close)
		fmt.Println(val)            // yeh received value print karega

		wg.Done() // WaitGroup ko batana ki yeh goroutine khatam ho gaya
	}(myCh, wg)

	// 👇 Ye goroutine sirf send karega channel mein
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0   // channel mein ek value bheji
		close(myCh) // channel ko close kar diya (ab koi value nahi bhej sakte)
		wg.Done()   // WaitGroup ko batana ki yeh goroutine bhi khatam ho gaya
	}(myCh, wg)

	wg.Wait() // jab tak dono goroutines complete nahi hoti, tab tak wait karo
}
