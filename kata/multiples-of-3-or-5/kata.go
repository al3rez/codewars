package kata

func Multiple3And5(number int) (sum int) {
	if number == 0 {
		return
	}

	if number%15 == 0 {
		return 1
	}

	for i := number - 1; i > 0; i-- {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return
}
