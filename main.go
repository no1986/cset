package main

import (
	"cpuset/cset"
	"fmt"
)

func main() {
	fmt.Println("main running ...")

	cset.CreateGroup("aa")
    cset.RenameGroup("aa", "bb")
	a, _ := cset.GroupList()
	for _, i := range a {
		fmt.Println(i)
	}
    b, _ := cset.ProcList("/")
    for _, j := range b {
        fmt.Println(j)
    }
	cset.DeleteGroup("bb")
    cset.DeleteGroup("aa")
}
