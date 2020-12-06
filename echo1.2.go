package main
import (
	"fmt"
	"os"
)
func main() {
	for A, arg := range os.Args[1:] {
		fmt.Println("Index no:",A ,"contains the value ",arg )
	}
}

