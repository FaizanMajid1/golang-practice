//Converts degree Fahrenheit To degree celcius .
package main
import (
	"fmt"
	"os"
	"strconv"
)
func main() {

    f, err := strconv.Atoi(os.Args[1])
    if err == nil {
        c := (f - 32) * 5 / 9
        fmt.Println("Temperature in Celcius is : ", c,"°C")
  	}
}