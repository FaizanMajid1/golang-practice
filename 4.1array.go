package main
import "fmt"

func main(){
	var myarray [3]int
	myarray = [...]int{20,25,30}
	fmt.Println(myarray[0]) //first element
	fmt.Println(myarray[len(myarray)-1]) //last element
	fmt.Println(myarray == [3]int{20,10,30} )
}