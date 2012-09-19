package fibo

var lastInput int

func Fibonacci(count int) int {
	lastInput = count

	if count <= 1 {
		return 1
	}
	return Fibonacci(count - 2) + Fibonacci(count - 1)
}
