package cset

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type CPUSetPath struct {
	BasePath string
	CPUs     string
	MEMs     string
	CPUX     string
	MEMX     string
	Tasks    string
}

type CPUSet struct {
	Group string
	CPUs  string
	MEMs  string
	CPUX  string
	MEMX  string
	Tasks int
}

type Proc struct {
	User   string
	PID    string
	PPID   string
	State  string
	CMD    string
}

func NewCPUSetPath() (*CPUSetPath, error) {
	csp := &CPUSetPath{}

	path, err := getBasePath()
	if err != nil {
		return nil, err
	}
	csp.BasePath = path
	csp.CPUs = "cpuset.cpus"
	csp.MEMs = "cpuset.mems"
	csp.CPUX = "cpuset.cpu_exclusive"
	csp.MEMX = "cpuset.mem_exclusive"
	csp.Tasks = "tasks"
	return csp, nil
}

func getBasePath() (string, error) {
	fp, err := os.Open("/proc/mounts")
	defer fp.Close()
	if err != nil {
		return "", err
	}

	scan := bufio.NewScanner(fp)
	var path string
	for scan.Scan() {
		line := scan.Text()
		n := strings.Contains(line, "cpuset")
		if n {
			path = strings.Fields(line)[1]
		}
	}
	return path, nil
}

func Write(path, str string) error {
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.WriteString(str)
	if err != nil {
		return err
	}
	return nil
}

func Read(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	str := strings.Split(string(b), "\n")[0]
	return str, nil
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

func (csp *CPUSetPath) GetTasks(group string) (int, []string, error) {
	path := filepath.Join(csp.BasePath, group, csp.Tasks)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, []string{}, err
	}
	lines := strings.Split(string(b), "\n")
	lines = lines[:len(lines)-1]
	return len(lines), lines, nil
}
