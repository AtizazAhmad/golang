package main
import (
	"fmt"
	"os"
	"strconv"
)
func main() {	
	var sum int
	for i:=1; i < len(os.Args); i++ {
		s:= os.Args[i]
	       n,err := strconv.Atoi(s)
	    if err == nil {
	    	sum+=n
	    }
	}
	fmt.Println(sum)
}
