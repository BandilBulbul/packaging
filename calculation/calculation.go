package calculation

func Add(values ...int) int {
	var sum int
	for _, value := range values {
		sum += value
	}
	return sum

}

func Mul(values ...int) int {
	var mul int
	for _, value := range values {
		mul *= value
	}
	return mul

}
