package main
import "fmt"
func main() {
x := 2
y := 8
p := &x         // p, of type *int, points to x
fmt.Println(y) // "8"
y = *p         // equivalent to y = x  i.e  y = 2
fmt.Println(y)  // "2"
fmt.Println(y == x, &x == p, x == *p) // "true true true"
}