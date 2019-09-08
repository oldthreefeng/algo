package main

import "fmt"

func gcd(a, b int) int {
	if a < 0 || b < 0 {
		return -1
	}
	if a%b != 0 {
		return gcd(b, a%b)
	}
	return b
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	a := gcd(35, 215)
	b := lcm(49, 35)
	fmt.Println(a, b) // 5 245
}
