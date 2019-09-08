package search

func Factorial(n int) int {
	if n < 2 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}

}
