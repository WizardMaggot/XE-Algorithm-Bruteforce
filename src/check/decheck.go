package check

import (
	"fmt"
	"os"
	"strings"
)

//sees if string is in text
func Check(fin *[]int, rot int, inpc string) {
	var chk []string
	var see string
	for c := 0; c < len(*fin); c++ {
		p := &(*fin)[c]
		chk = append(chk, string(rune(*p-rot)))
		see = strings.Join(chk, "")
	}
	if strings.Contains(see, inpc) {
		println("Found!\n" + "l")
		os.Exit(0)
	} else {
		//lock.Done()
		fmt.Println("HiHi!")
	}
}
