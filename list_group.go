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
	}

	cset.DeleteGroup("system")
}
