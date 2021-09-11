package decode

import (
	"fmt"

	"github.com/WizardMaggot/XE-Algorithm-Bruteforce/check"
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
func Generate(d []int) func(string) {
	g, h := findMinAndMax(d)
	return func(ip string) {
		for i := h - 255; i < g+1; i++ {
			check.Check(d, i, ip)
			//lock.Add(1)
			//go ck(i, ip)
		}
	}
}
