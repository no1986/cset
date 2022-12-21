package cset

import (
	"os"
	"path/filepath"
	"strings"
)

type Group struct {
	Name  string
	CPUs  string
	MEMs  string
	CPUx  string
	MEMx  string
	Tasks int
}

func GroupList() ([]Group, error) {
	ret := []Group{}
	p, err := GetPath()
	if err != nil {
		return ret, err
	}
	dirs, err := WalkDir(p.base)
	if err != nil {
		return ret, err
	}

	for _, dir := range dirs {
		group := dir[len(p.base):]
		if group == "" {
			group = "/"
		}

		g := new(Group)
		g.Name = group

		path := filepath.Join(p.base, group, p.cpus)
		cpus, err := GetPathVal(path)
		if err != nil {
			return ret, err
		}
		if len(cpus) > 0 {
			g.CPUs = cpus[0]
		}

		path = filepath.Join(p.base, group, p.mems)
		mems, err := GetPathVal(path)
		if err != nil {
			return ret, err
		}
		if len(mems) > 0 {
			g.MEMs = mems[0]
		}

		path = filepath.Join(p.base, group, p.cpux)
		cpux, err := GetPathVal(path)
		if err != nil {
			return ret, err
		}
		g.CPUx = cpux[0]

		path = filepath.Join(p.base, group, p.memx)
		memx, err := GetPathVal(path)
		if err != nil {
			return ret, err
		}
		g.MEMx = memx[0]

		path = filepath.Join(p.base, group, p.tasks)
		tasks, err := GetPathVal(path)
		if err != nil {
			return ret, err
		}
		g.Tasks = len(tasks)
		ret = append(ret, *g)
	}
	return ret, nil
}

func CreateGroup(group string) error {
	p, err := GetPath()
	if err != nil {
		return err
	}
	path := filepath.Join(p.base, group)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// 既にグループが存在する場合は処理を正常終了
		return nil
	}

	if err := os.Mkdir(path, 0777); err != nil {
		return err
	}

	tmp := strings.Split(group, "/")
	parent := strings.Join(tmp[:len(tmp)-1], "/")

	path = filepath.Join(p.base, parent, p.cpus)
	cpus, err := GetPathVal(path)
	if err != nil {
		return err
	}
	if err = SetCPUs(group, cpus[0]); err != nil {
		return err
	}

	path = filepath.Join(p.base, parent, p.mems)
	mems, err := GetPathVal(path)
	if err != nil {
		return err
	}
	if err = SetMEMs(group, mems[0]); err != nil {
		return err
	}

	return nil
}

func RenameGroup(old, new string) error {
	p, err := GetPath()
	if err != nil {
		return err
	}
	oldpath := filepath.Join(p.base, old)
	newpath := filepath.Join(p.base, new)
	if err := os.Rename(oldpath, newpath); err != nil {
		return err
	}
	return nil
}

func DeleteGroup(group string) error {
	p, err := GetPath()
	if err != nil {
		return err
	}
	path := filepath.Join(p.base, group)
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}

func SetCPUs(group, val string) error {
	p, err := GetPath()
	if err != nil {
		return err
	}

	path := filepath.Join(p.base, group, p.cpus)
	if err = SetPathVal(path, val); err != nil {
		return err
	}
	return nil
}

func SetMEMs(group, val string) error {
	p, err := GetPath()
	if err != nil {
		return err
	}

	path := filepath.Join(p.base, group, p.mems)
	if err = SetPathVal(path, val); err != nil {
		return err
	}
	return nil
}

func SetCPUX(group, val string) error {
	p, err := GetPath()
	if err != nil {
		return err
	}

	path := filepath.Join(p.base, group, p.cpux)
	if err = SetPathVal(path, val); err != nil {
		return err
	}
	return nil
}

func SetMEMX(group, val string) error {
	p, err := GetPath()
	if err != nil {
		return err
	}

	path := filepath.Join(p.base, group, p.memx)
	if err = SetPathVal(path, val); err != nil {
		return err
	}
	return nil
}
