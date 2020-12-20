package greet
import "fmt"

var Greeting string
Greeting = "AssalamAliqum"

func Hello(name string) string {
	return fmt.Println(Greeting, name)
}