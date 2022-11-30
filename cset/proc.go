package cset

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

func (csp *CPUSetPath) ProcList(group string) ([]Proc, error) {
	ret := []Proc{}

	_, tasks, err := csp.GetTasks(group)
	if err != nil {
		return ret, err
	}

	for _, pid := range tasks {
		proc := new(Proc)

		path := filepath.Join("/proc", pid, "cmdline")
		cmd, err := Read(path)
		if err != nil {
			return ret, nil
		}
		proc.CMD = cmd

		path = filepath.Join("/proc", pid, "status")
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return ret, err
		}
		lines := strings.Split(string(b), "\n")
		lines = lines[:len(lines)-1]
		for _, line := range lines {
			t := strings.Split(line, ":")
			key := t[0]
			val := strings.TrimSpace(t[1])
			if key == "Uid" {
				val = strings.Fields(val)[0]
				usr, err := user.LookupId(val)
				if err != nil {
					return ret, err
				}
				proc.User = usr.Username

			} else if key == "Pid" {
				proc.PID = val
			} else if key == "PPid" {
				proc.PPID = val
			} else if key == "State" {
				proc.State = strings.Fields(val)[0]
			} else if key == "Name" && proc.CMD == "" {
				proc.CMD = fmt.Sprintf("[%v]", val)
			}
		}

		ret = append(ret, *proc)
	}
	return ret, nil
}
