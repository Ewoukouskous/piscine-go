package quete8

func Any(f func(string) bool, a []string) bool {
	test := false
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			test = true
		}
	}
	return test
}
