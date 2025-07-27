package main
import "fmt"
//methods
//function in classes is known as methods
func main() {
	fmt.Println("Structs in GoLang")
	hitesh := User{"hitesh", "h@gmail.com", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh detials are:%+v\n", hitesh)
	fmt.Printf("hitesh detials are:%+v\n", hitesh)
	fmt.Printf("name is %v and email is %v", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()
}
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
func (u User) GetStatus(){
	fmt.Println("is user active:",u.Status)
}
func(u User) NewMail(){
	u.Email="test@go.dev"
	fmt.Println("email is",u.Email)
}
//it will not replace the original,just manipulate
