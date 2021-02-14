package main
import(
	"fmt"
	"github.com/FaizanMajid1/go-packages/tempconv"
)
func main(){
var tempC tempconv.Celsius
var tempF tempconv.Fahrenheit

tempF = 32.0
tempC = 100.0

fmt.Println(tempconv.FtoC(tempF), "=", tempF)
fmt.Println(tempC, "=", tempconv.CtoF(tempC))


}

