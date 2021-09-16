package primer

import (
	"strconv"
	"strings"

	"github.com/WizardMaggot/XE-Algorithm-Bruteforce/pirate"
)

var fin []int
var conj []int

//every 3 numbers become one in an array
func Prime(st string) []int {
	//remove all new lines
	st = strings.ReplaceAll(st, "\n", "")
	st = strings.ReplaceAll(st, " ", "")
	ch := strings.Split(st, ".")

	//individual numbers to array
	for _, i := range ch {
		if i != "" {
			n, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			conj = append(conj, n)
		}
	}
	//add every 3 numbers
	for dc := 0; dc < len(conj)/3; dc++ {
		//faster to pass slices
		fin = append(fin, pirate.Sumarr(conj[0+(3*dc):3+(3*dc)]))
	}
	return fin
}
