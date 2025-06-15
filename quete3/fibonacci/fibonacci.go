package piscine

func Fibonacci(arg int) int {
	if arg < 0 {
		return -1
	}
	if arg == 0 {
		return 0
	} else if arg == 1 {
		return 1
	} else {
		return Fibonacci(arg-1) + Fibonacci(arg-2)
	}
}
