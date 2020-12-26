//Converts degree Fahrenheit To degree celcius .
package main
import "fmt"
     func main() {
     	var f int
     	fmt.Println("Enter temperature in Fahrenheit : ")
     	fmt.Scanln(&f)
        var c = (f - 32) * 5 / 9
        fmt.Println("Temperature in Celcius is : ", c,"°C")
  
}