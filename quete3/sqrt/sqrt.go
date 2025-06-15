package piscine

func Sqrt(n int) int {
	if n == 1 {
		return 1
	}
	for i := 1; i < n; i++ {
		if i*i == n {
			return i
		}
	}
	return 0
}
