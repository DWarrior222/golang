package main
import (
	"fmt"
	"strings"
)
func makeSuffix(suffix string) func (string) string {

	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	suffix := makeSuffix(".jpg")
	fmt.Println("hello.jpg", suffix("hello.jpg"))
	fmt.Println("hello.j", suffix("hello.j"))

}
