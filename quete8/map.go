package quete8

func Map(f func(int) bool, a []int) []bool {
	slice := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		slice[i] = f(a[i])
	}
	return slice
}
