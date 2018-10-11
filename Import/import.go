package main
import "fmt"
import "math"
import M "math"
import . "math"

func main() {
   a, b := 345.22, 300.56
   fmt.Println("Minimum num", math.Min(a, b))
   fmt.Println("Minimum num", M.Min(a, b))
   fmt.Println("Minimum num", Min(a, b))
}
