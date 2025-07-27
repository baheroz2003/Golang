package main

import "fmt"

func main() {
	fmt.Println("maps in golang")
	languages:=make(map[string]string)
	languages["r"]="ruby"
	languages["c"]="cpp"
	fmt.Println("list all",languages)
	fmt.Println("r is",languages["r"])
	delete(languages,"r")
	fmt.Println("list all",languages)
	//key value pair
    //loop through the map
	//:= is called vulrous operator
	// for key,value:=range languages{
	// 	fmt.Printf("for key %v,value is %v\n",key,value)
	// }
	for _,value:=range languages{
		fmt.Printf("for key v,value is %v\n",value)
	}
	//if only want value to ignore key
}
