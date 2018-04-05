package kata

func FindUniq(arr []float32) float32 {
	h := map[float32]int{}

	for _, f := range arr {
		h[f] += 1
	}

	for f, _ := range h {
		if h[f] == 1 {
			return f
		}
	}

	return 0
}
