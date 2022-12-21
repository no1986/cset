package cset

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Proc struct {
	User   string
	Pid    string
	PPid   string
	State  string
	CMD    string
	Exe    string
	CPUs   string
	MEMs   string
	bound  bool
	kernel bool
}

func ProcList(group string) ([]Proc, error) {
	ret := []Proc{}
	p, err := GetPath()
	if err != nil {
		return ret, err
	}

	path := filepath.Join(p.base, group, "tasks")
	tasks, err := GetPathVal(path)
	if err != nil {
		return ret, err
	}

	root_cpus, err := GetPathVal(filepath.Join(p.base, p.cpus))
	if err != nil {
		return ret, err
	}
	root_mems, err := GetPathVal(filepath.Join(p.base, p.mems))
	if err != nil {
		return ret, err
	}

	for _, pid := range tasks {
		proc := new(Proc)

		path = filepath.Join("/proc", pid, "cmdline")
		cmd, err := GetPathVal(path)
		if err != nil {
			return ret, err
		}

		path = filepath.Join("/proc", pid, "status")
		status, err := GetPathVal(path)
		if err != nil {
			return ret, err
		}
		statusd := map[string]string{}
		for _, s := range status {
			ss := strings.Fields(s)
			if len(ss) <= 1 {
				continue
			}
			key := ss[0]
			key = strings.TrimSpace(key[:len(key)-1])
			val := strings.TrimSpace(string(ss[1]))
			statusd[key] = val
		}

		proc.Pid = pid
		path := filepath.Join("/proc", pid, "exe")
		info, err := os.Lstat(path)
		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			if link, err := os.Readlink(path); err == nil {
				proc.Exe = link
				proc.kernel = false
			} else {
				proc.kernel = true
			}
		} else {
			proc.kernel = true
		}

		if len(cmd) > 0 {
			proc.CMD = strings.ReplaceAll(cmd[0], string([]byte{0}), " ")
		} else {
			proc.CMD = fmt.Sprintf("[%s]", statusd["Name"])
		}
		proc.PPid = statusd["PPid"]
		proc.State = statusd["State"]
		u, err := user.LookupId(statusd["Uid"])
		if err != nil {
			return ret, err
		}
		proc.User = u.Username
		proc.CPUs = statusd["Cpus_allowed_list"]
		proc.MEMs = statusd["Mems_allowed_list"]
		if proc.CPUs == root_cpus[0] && proc.MEMs == root_mems[0] {
			proc.bound = false
		} else {
			proc.bound = true
		}

		ret = append(ret, *proc)
	}
	return ret, nil
}
