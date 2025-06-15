package piscine

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(n int) int {
	if IsPrime(n) {
		return n
	} else {
		for i := n; i >= n; i++ {
			if IsPrime(i) {
				return i
			}
		}
	}
	return 3673
}
