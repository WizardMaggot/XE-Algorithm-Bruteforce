package pirate

//Minimum of an array
func Min(values []int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

//Minimum of an array
func Max(values []int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

//finds the sum of an array, x is a pointer
func Sumarr(array []int) int {
	result := array[0] + array[1] + array[2]
	return result
}
