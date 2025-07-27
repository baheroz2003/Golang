package main
import "fmt"
func main() {
	fmt.Println("loops in golang")
    days:=[]string{"sunday","tuesday","wednesday"}
	fmt.Println(days)
	for d:=0;d<len(days);d++{
		fmt.Println(days[d])
	}
	for i:=range days{
		fmt.Println(days[i])
	}
	for idx,day:=range days{
		fmt.Printf("index is %v and value is %v\n",idx,day)
	}
	rogueval:=1
	for rogueval<10{
		if(rogueval==2){
			goto lco
		}
		if(rogueval==5){
			rogueval++
			continue
		}
		fmt.Println("value is:",rogueval)
		rogueval++;
	}
	lco:
		fmt.Println("jump at learncoder")

}
