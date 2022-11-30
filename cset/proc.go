package cset

import (
	"fmt"
	"path/filepath"
)

func (csp *CPUSetPath) ProcList(group string) error {
	_, tasks, err := csp.GetTasks(group)
	if err != nil {
		return err
	}
	for _, pid := range tasks {
		fmt.Println(pid)
		path := filepath.Join("/proc", pid)
	}
	return nil
}
