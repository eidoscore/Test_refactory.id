// for j := 0; j >= 1; j-- {
// 	fmt.Println( j)
// }

package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 6; i >= 1; i-- {
		fmt.Printf("%s\n", strings.Repeat("#", i))
	}
}
