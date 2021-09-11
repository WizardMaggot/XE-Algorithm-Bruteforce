package primer

import (
	"strconv"
	"strings"
)

//every 3 numbers become one in an array
func Prime(st string) []int {
	var temp []int
	var dc int

	var fin []int

	//every line
	ln := strings.Split(st, "\n")
	for l := range ln {
		//every .
		ch := strings.Split(ln[l], ".")
		//every
		for c := range ch {
			if ch[c] != "" {
				nu, _ := strconv.Atoi(ch[c])
				temp = append(temp, nu)
				dc += 1
				if dc%3 == 0 {
					fin = append(fin, sum(temp))
					temp = nil
				}
			}
		}
	}
	return fin
}

//finds the sum of an array
func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
