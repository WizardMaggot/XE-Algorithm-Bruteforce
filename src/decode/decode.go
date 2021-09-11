package decode

import (
	"fmt"
)

//Minimum and max of an array
func findMinAndMax(a []int) (min int, max int) {
	/*for _, v := range a {
		if v < min {
			min = v
		}
	}*/
	fmt.Println(a)
	return a[0], a[1]
}

//bruteforces text and check() if string is inside
func Generate(d []int) (int, int) {
	g, h := findMinAndMax(d)
	return g, h

	/*for i := h - 255; i < g+1; i++ {
	lock.Add(1)
	go check(i, ip)*/
}
