package cset

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Path struct {
	base  string
	cpus  string
	mems  string
	cpux  string
	memx  string
	tasks string
}

var path *Path

func GetPath() (*Path, error) {
	if path == nil {
		fp, err := os.Open("/proc/mounts")
		if err != nil {
			return nil, err
		}
		defer fp.Close()

		scan := bufio.NewScanner(fp)
		var base string
		for scan.Scan() {
			line := scan.Text()
			n := strings.Contains(line, "cpuset")
			if n {
				base = strings.Fields(line)[1]
			}
		}
		path = &Path{
			base:  base,
			cpus:  "cpuset.cpus",
			mems:  "cpuset.mems",
			cpux:  "cpuset.cpu_exclusive",
			memx:  "cpuset.mem_exclusive",
			tasks: "tasks",
		}
	}
	return path, nil
}

func GetPathVal(path string) ([]string, error) {
	var ret []string
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return ret, err
	}
	for _, s := range strings.Split(string(b), "\n") {
		if s != "" {
			ret = append(ret, s)
		}
	}
	return ret, nil
}

func SetPathVal(path, val string) error {
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	if _, err = fp.WriteString(val); err != nil {
		return err
	}
	return nil
}

func WalkDir(path string) ([]string, error) {
	dirs := []string{}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		dirs = append(dirs, path)
		return nil
	})
	return dirs, nil
}
