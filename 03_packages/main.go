package main

import (
	"fmt"
	"math"
	"github.com/sirghostal/go_hello/03_packages/strutil"
	)

func main() {
	// math.Floor rounds whatever value you provide DOWN.
	fmt.Println(math.Floor(9.6))
	// math.Ceil rounds UP.
	fmt.Println(math.Ceil(9.6))
	// math.Sqrt .. squareroots.
	fmt.Println(math.Sqrt(64))
	fmt.Println(strutil.Reverse("Hello World!"))
}