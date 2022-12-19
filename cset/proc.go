package cset

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Proc struct {
	user  string
	pid   string
	ppid  string
	state string
	cmd   string
	exe   string
	cpus  string
	mems  string
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

		proc.pid = pid
		path := filepath.Join("/proc", pid, "exe")
		info, err := os.Lstat(path)
		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			link, err := os.Readlink(path)
			if err == nil {
				proc.exe = link
			}
		}
		if len(cmd) > 0 {
			proc.cmd = strings.ReplaceAll(cmd[0], string([]byte{0}), " ")
		} else {
			proc.cmd = fmt.Sprintf("[%s]", statusd["Name"])
		}
		proc.ppid = statusd["PPid"]
		proc.state = statusd["State"]
		u, err := user.LookupId(statusd["Uid"])
		if err != nil {
			return ret, err
		}
		proc.user = u.Username
		proc.cpus = statusd["Cpus_allowed_list"]
		proc.mems = statusd["Mems_allowed_list"]

		ret = append(ret, *proc)
	}
	return ret, nil
}
