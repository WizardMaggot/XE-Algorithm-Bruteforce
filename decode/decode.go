package decode

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

//make sure everything is gone through
var lock sync.WaitGroup

//bruteforces text and check() if string is inside
func (d []int) Decode(ip string) {
	g, h := findMinAndMax(d)
	fmt.Println(g, h)
	for i := h - 255; i < g+1; i++ {
		lock.Add(1)
		go check(i, ip)
	}
}

//sees if string is in text
func check(u int, inpc string) {
	var chk []string
	var see string
	for c := 0; c < len(fin); c++ {
		chk = append(chk, string(rune(fin[u])))
		see = strings.Join(chk, "")
	}
	if strings.Contains(see, inpc) {
		println("Found!\n" + "l")
		os.Exit(0)
	} else {
		lock.Done()
	}
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
