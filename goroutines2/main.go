package main
import (
	"fmt"
	"sync"
)
func main() {
	fmt.Println("race condition - LearnCodeOnline.in")
	wg := &sync.WaitGroup{}  // WaitGroup banaya goroutines ke complete hone ka wait karne ke liye
	mut := &sync.RWMutex{}   // RWMutex banaya slice ke access ko synchronize karne ke liye
	var score = []int{0}     // Ek slice banaya jo sab goroutines modify karenge

	wg.Add(3) // 3 goroutines chalenge, toh WaitGroup me 3 Add kara

	// 1st goroutine
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("one r")       // Print karega
		m.Lock()                   // Exclusive write access ke liye Lock
		score = append(score, 1)   // Slice me 1 add karega
		m.Unlock()                 // Lock release karega
		wg.Done()                  // Work complete, ek goroutine done
	}(wg, mut)

	// 2nd goroutine
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("two r")
		m.Lock()
		score = append(score, 2)   // Slice me 2 add karega
		m.Unlock()
		wg.Done()
	}(wg, mut)

	// 3rd goroutine
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("three r")
		m.Lock()
		score = append(score, 3)   // Slice me 3 add karega
		m.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()                   // Sabhi goroutines ke complete hone ka wait karega

	fmt.Println(score)         // Final result print karega
}
// sync.Mutex--Agar ek banda resource use kar raha hai (read/write), toh doosre sab wait karenge, chahe unhe sirf read hi karna ho.
//sync.RWMutex--Agar sirf read karna hai, toh kai log ek saath padh sakte hain. Par agar koi likhne aaya (write), toh sabka access roka jayega.
