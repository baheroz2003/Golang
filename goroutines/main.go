package main // yeh main package hai, program execution yahin se start hoga

import (
	"fmt"      // print karne ke liye
	"net/http" // HTTP requests bhejne ke liye (websites check karne ke liye)
	"sync"     // goroutines ke sync karne ke liye (WaitGroup ka use hoga)
)

// var wg sync.WaitGroup //pointer
var wg sync.WaitGroup // WaitGroup ek sync tool hai jo multiple goroutines complete hone ka wait karta hai

func main() {
	// go greeter("hello") //background me chlega hello
	// greeter("world")    //0-3 tk world jitne baar
	//fir hello 3-6
	//fir world 6
	//6-9 me hello hello
	//world world
	//world main thread pe chl rha hai

	// 👆 yeh comments purane example ke liye the (greeter function) — batate hain ki kaise goroutines background mein chalte hain aur kaise output mix ho jata hai.

	websiteList := []string{ // ek list banayi websites ki jise check karna hai
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList { // list ke har element (website) par loop
		go getStatusCode(web) // har website ke liye background mein ek goroutine chalaya

		//if u are using you should wait there 
		//use time concept
		//use sync method

		// 👆 yeh kehna chah rahe ho ki agar goroutines use kar rahe ho to rukna padega
		// WaitGroup se sync karo ya time.Sleep mat karo (WaitGroup best practice hai)

		wg.Add(1) // har new goroutine ke liye WaitGroup mein +1 (bataya ki ek aur kaam start hua)
		//keep on adding no of go routines 
	}
	wg.Wait() //please don't exit wait my friend are coming
	// 👆 jab tak sabhi goroutines khatam na ho jayein, main function rukega
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3*time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
// 👆 yeh purana function tha background mein string print karne ka, ab is code mein use nahi ho raha

func getStatusCode(endpoint string) {
	defer wg.Done() //at last
	// 👆 jab function complete hoga, WaitGroup mein -1 ho jayega (kaam complete)

	res, err := http.Get(endpoint) // yahan GET request bhej rahe hain website par
	if err != nil {
		fmt.Println("oops in endpoint") // agar koi error aaye to print karo
	} else {
		fmt.Printf(" %d 200 status code for %s\n", res.StatusCode, endpoint)
		// agar request successful ho to website ka status code print karo
	}
}
