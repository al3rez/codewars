package kata

func Solution(str string) (final []string) {
	if len(str)%2 != 0 {
		str = str + "_"
	}
	var r []byte
	for i := 0; i < len(str); i++ {
		r = append(r, str[i])
		if len(r) == 2 {
			final = append(final, string(r))
			r = []byte{}
		}
	}
	return
}
