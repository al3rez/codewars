package kata

func multipleOfIndex(ints []int) (m []int) {
	for i, n := range ints {
		if i != 0 && n%i == 0 {
			m = append(m, n)
		}
	}
	return
}
