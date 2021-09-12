package primer

import (
	"strconv"
	"strings"
)

var fin []int
var conj []int

//every 3 numbers become one in an array
func Prime(st string) []int {
	//remove all new lines
	st = strings.ReplaceAll(st, "\r\n", "")
	st = strings.ReplaceAll(st, " ", "")
	ch := strings.Split(st, ".")
	//individual numbers to array
	for _, i := range ch {
		if i != "" {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			conj = append(conj, j)
		}
	}
	//add every 3 numbers
	z := len(conj) / 3
	var p int
	for dc := 0; dc < z; dc++ {
		fin = append(fin, sum(conj, p))
		p += 1
	}
	return fin
}

//finds the sum of an array
func sum(array []int, x int) int {
	result := array[0+3*x] + array[1+3*x] + array[2+3*x]
	return result
}
