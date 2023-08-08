//go:build linux

/*

Original Repository
https://github.com/mitchellh/go-ps

This version is just for personal usages.. "wink wink"
So, No support for this

*/

package process

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type UnixProcess struct {
	pid    int
	binary string
}

func (p *UnixProcess) Pid() int {
	return p.pid
}

func (p *UnixProcess) Executable() string {
	return p.binary
}

func processes() ([]Process, error) {
	d, err := os.Open("/proc")
	if err != nil {
		return nil, err
	}
	defer d.Close()

	results := make([]Process, 0, 50)
	for {
		names, err := d.Readdirnames(10)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for _, name := range names {
			if name[0] < '0' || name[0] > '9' {
				continue
			}
			pid, err := strconv.ParseInt(name, 10, 0)
			if err != nil {
				continue
			}
			p, err := newUnixProcess(int(pid))
			if err != nil {
				continue
			}
			results = append(results, p)
		}
	}
	return results, nil
}

func newUnixProcess(pid int) (*UnixProcess, error) {
	p := &UnixProcess{pid: pid}
	return p, p.Refresh()
}

func (p *UnixProcess) Refresh() error {
	statPath := fmt.Sprintf("/proc/%d/cmdline", p.pid)
	dataBytes, err := ioutil.ReadFile(statPath)
	if err != nil {
		return err
	}

	p.binary = strings.Replace(string(dataBytes), "\x00", " ", -1)
	return nil
}
