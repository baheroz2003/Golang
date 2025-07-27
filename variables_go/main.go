package main
import "fmt"
//not allowed
// jwTtOKEN :=300000
// var jwttoken int=300000
//allowed
const LoginToken string = "ghabbhhjd" //public
func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)
	var islog bool = false
	fmt.Println(islog)
	fmt.Printf("variable is of type: %T \n", islog)
	var smallflat float32 = 255.6
	fmt.Println(smallflat)
	fmt.Printf("variable is of type: %T \n", smallflat)
	//default value and some classes
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type: %T \n", anothervariable)
	//implicit type
	//type should be defined
	var website = "learncodeonline.in"
	fmt.Println(website)
	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Println(LoginToken)
}
