package main
import (
	"fmt"
	"os"
)
func main() {
	for I, arg := range os.Args[1:] {
		fmt.Println("Index",I ,"contains the value ",arg )
	}
}