package kata

func century(year int) (c int) {
	if year <= 100 {
		return 1
	}

	c = year / 100
	if (year % 100) > 0 {
		c++
	}

	return
}
