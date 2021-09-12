package main

import (
	"fmt"
	"strconv"
)

func main() {
	var t = []string{"1", "2", "3"}
	var t2 []int

	for _, i := range t {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	fmt.Println(t2)
}
