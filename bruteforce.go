package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/WizardMaggot/XE-Algorithm-Bruteforce/primer"
)

//final "primed" string
var fin []int

//make sure everything is gone through
var lock sync.WaitGroup

//basic input and code exec
func main() {
	//var input string
	input := "ex.txt"
	//fmt.Printf("Enter filename> ")
	//fmt.Scanln(&input)
	if input == "" {
		fmt.Println("Enter filename\n> ")
		main()
	} else {
		f, err := os.ReadFile(input)
		if err != nil {
			fmt.Println("Failed to read file")
			main()
		}
		cf := string(f)
		fin = primer.Prime(cf)
		fmt.Printf("Key letters?\n> ")
		fmt.Scanln(&input)

		fmt.Println("fin")

		decode(input)
	}
}

//bruteforces text and check() if string is inside
func decode(ip string) {
	g, h := findMinAndMax(fin)
	fmt.Println(g, h)
	for i := h - 255; i < g+1; i++ {
		lock.Add(1)
		go check(i, ip)
	}
	defer lock.Wait()
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
