// dice game like in ludo
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch and case in golang")
	rand.Seed(time.Now().UnixNano())

	dicenumber := rand.Intn(6) + 1
	fmt.Println("value of dice is", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("you can open")
	case 2:
		fmt.Println("you can move 2 spaces")
	case 3:
		fmt.Println("you can move 3 spaces")
		fallthrough
	case 4:
		fmt.Println("you can move 4 spaces")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spaces")
	case 6:
		fmt.Println("you can move to 6 spot and roll dice again")
	default:
		fmt.Println("what was that!")
	} // <-- this was missing
}
