package piscine

func SortIntegerTable(table []int) []int {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[i] < table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
	return table
}
