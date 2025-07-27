package main
import "fmt"
func main() {
	fmt.Println("Structs in GoLang")
	hitesh := User{"hitesh", "h@gmail.com", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh detials are:%+v\n", hitesh)
	fmt.Printf("hitesh detials are:%+v\n", hitesh)
	fmt.Printf("name is %v and email is %v", hitesh.Name, hitesh.Email)
}
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
