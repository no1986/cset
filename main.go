package main

import (
	"cpuset/cset"
	"fmt"
)

func main() {
	fmt.Println("main running ...")

	fmt.Println()
	csp, _ := cset.NewCPUSetPath()
	groups, _ := csp.GroupList()
	for _, g := range groups {
		fmt.Printf("%+v\n", g)
	}
	fmt.Println()

	//err := csp.ProcList("/docker/98aa602495f217ded9d9928745b976aed228ff081ab1c185b5489ca25837e15c")
	ret, _ := csp.ProcList("/kubepods/burstable/pod0debeaf2-fcd8-48ae-a4cc-0fa9b8272843/3031a428d7e64f6935bb93e5fdf50f2325a3f83ddc1b5d0660b9504c47a251e4")
	//err := csp.ProcList("/")
	fmt.Printf("%+v\n", ret)
}
