// defer invokes a fucntion whose exception is deferred to the moment the surrounding function returns,either because the surrounding fucntion executed a return statement reaches the end of its function body or because the corresponding goroutine is panicking
// defer lgao toh wo last me execute hoga
// last in first out
package main

import "fmt"

func main() {
	 defer fmt.Println("hello")
	  defer fmt.Println("world")
	  //in reverse order they print
	  mydefer()
}
func mydefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}
