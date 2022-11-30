package cset

import (
	"os"
	"path/filepath"
	"strings"
)

func (csp *CPUSetPath) GroupList() ([]CPUSet, error) {
	dirs, err := WalkDir(csp.BasePath)
	ret := []CPUSet{}
	if err != nil {
		return ret, err
	}

	for _, dir := range dirs {
		s := new(CPUSet)
		group := dir[len(csp.BasePath):]
		if group == "" {
			group = "/"
		}
		s.Group = group
		cpus, err := csp.Get(group, csp.CPUs)
		if err != nil {
			return ret, err
		}
		s.CPUs = cpus
		mems, err := csp.Get(group, csp.MEMs)
		if err != nil {
			return ret, err
		}
		s.MEMs = mems
		cpuX, err := csp.Get(group, csp.CPUX)
		if err != nil {
			return ret, err
		}
		s.CPUX = cpuX
		memX, err := csp.Get(group, csp.MEMX)
		if err != nil {
			return ret, err
		}
		s.MEMX = memX
		tasks, _, err := csp.GetTasks(group)
		if err != nil {
			return ret, err
		}
		s.Tasks = tasks
		ret = append(ret, *s)
	}
	return ret, nil
}

func (csp *CPUSetPath) Create(group string) error {
	path := filepath.Join(csp.BasePath, group)
	if f, err := os.Stat(path); !os.IsNotExist(err) && f.IsDir() {
		// 既にグループが存在する場合は処理を正常終了
		return nil
	}
	if err := os.Mkdir(path, 0777); err != nil {
		return err
	}

	t := strings.Split(group, "/")
	parent := strings.Join(t[:len(t)-1], "/")

	cpus, err := csp.Get(parent, csp.CPUs)
	if err != nil {
		return err
	}
	err = csp.Set(group, cpus, csp.CPUs)
	if err != nil {
		return err
	}

	mems, err := csp.Get(parent, csp.MEMs)
	if err != nil {
		return err
	}
	err = csp.Set(group, mems, csp.MEMs)
	if err != nil {
		return err
	}
	return nil
}

func (csp *CPUSetPath) Delete(group string) error {
	path := filepath.Join(csp.BasePath, group)
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}

func (csp *CPUSetPath) Rename(oldGroup, newGroup string) error {
	oldPath := filepath.Join(csp.BasePath, oldGroup)
	newPath := filepath.Join(csp.BasePath, newGroup)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}
	return nil
}

func (csp *CPUSetPath) Set(group, str, file string) error {
	path := filepath.Join(csp.BasePath, group, file)
	err := Write(path, str)
	if err != nil {
		return err
	}
	return nil
}

func (csp *CPUSetPath) Get(group, file string) (string, error) {
	path := filepath.Join(csp.BasePath, group, file)
	cpus, err := Read(path)
	if err != nil {
		return "", err
	}
	return cpus, nil
}
