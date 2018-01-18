package templates

func Until(count int) []int {
	step := 1
	if count < 0 {
		step = -1
	}
	return untilStep(0, count, step)
}
func untilStep(start, stop, step int) []int {
	v := []int{}

	if stop < start {
		if step >= 0 {
			return v
		}
		for i := start; i > stop; i += step {
			v = append(v, i)
		}
		return v
	}

	if step <= 0 {
		return v
	}
	for i := start; i <= stop; i += step {
		v = append(v, i)
	}
	return v
}
