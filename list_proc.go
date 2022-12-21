package main

import (
	"cpuset/cset"
	"fmt"
)

func main() {
	fmt.Println("main running ...")

	a, _ := cset.ProcList("/")
	for _, i := range a {
		fmt.Printf("%+v\n", i)
	}
}
