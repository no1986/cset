package main

import (
	"cpuset/cset"
	"fmt"
)

func main() {
	fmt.Println("main running ...")

	cset.CreateGroup("system")
	a, _ := cset.GroupList()
	for _, i := range a {
		fmt.Printf("%+v\n", i)
		b, _ := cset.ProcList(i.Name)
		for _, j := range b {
			fmt.Printf("%+v\n", j)
		}
		fmt.Println("\n")
	}

	cset.DeleteGroup("system")
}
