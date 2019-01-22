package main
// import "fmt"
// import "rand"
// import "time"
import (
    "fmt"
    "math/rand"
    "time"
)
func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("n", rand.Intn(100) + 1)
}
