package piscine

func IterativeFactorial(n int) int {
	res := 1
	if n > 21 || n < -21 {
		return 0
	}
	if n == 0 {
		return 1
	}
	for i := n; i > 0; i-- {
		res *= i
	}
	return res
}
