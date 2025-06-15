package piscine

func RecursiveFactorial(n int) int {
	res := 0
	if n < 21 && n > -21 {
		if n < 2 {
			return 1
		} else {
			newN := n - 1
			res += n * RecursiveFactorial(newN)
		}
	}
	return res
}
