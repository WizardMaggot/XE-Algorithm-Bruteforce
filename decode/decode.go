package decode

import (
	"fmt"
)

//make sure everything is gone through

//bruteforces text and check() if string is inside
func (d *[]int) Decode() (int, int) {
	g, h := findMinAndMax(*d)
	fmt.Println(g, h)
	return g, h
	/*
		for i := h - 255; i < g+1; i++ {
			lock.Add(1)
			go check(i, ip)
		}*/
}

//Minimum and max of an array
func findMinAndMax(a []int) (min int, max int) {
	fmt.Println(a[0])
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return a[0], a[1]
}
