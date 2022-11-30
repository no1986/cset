package main

import (
	"cpuset/cset"
	"fmt"
)

func main() {
	fmt.Println("main running ...")

	fmt.Println()
	cpuset, _ := cset.NewCPUSetPath()
	groups, _ := cpuset.GroupList()
	for _, g := range groups {
		fmt.Printf("%+v\n", g)
	}
	fmt.Println()

	cpuset.ProcList("/")
}
