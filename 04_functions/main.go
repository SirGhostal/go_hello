package main

import "fmt"

func greet(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}
func main() {
	fmt.Println(greet("Bob"))
	fmt.Println(getSum(5, 2))
}
