//go:build windows

/*

Original Repository
https://github.com/mitchellh/go-ps

This version is just for personal usages.. "wink wink"
So, No support for this

*/

package process

import (
	"os/exec"
	"strconv"
	"strings"
)

type WindowsProcess struct {
	pid int
	exe string
}

func (p *WindowsProcess) Pid() int {
	return p.pid
}

func (p *WindowsProcess) Executable() string {
	return p.exe
}

func newWindowsProcess(pid int, exe string) *WindowsProcess {
	return &WindowsProcess{
		pid: pid,
		exe: exe,
	}
}

func processes() ([]Process, error) {
	out, err := exec.Command("wmic", "process", "get", "ProcessId,CommandLine", "/format:list").Output()
	if err != nil {
		return nil, err
	}

	lines := strings.FieldsFunc(string(out), func(r rune) bool {
		return r == '\r' || r == '\n'
	})

	var processes []Process
	for _, line := range lines {
		fields := strings.SplitN(line, "=", 2)
		if len(fields) != 2 {
			continue
		}

		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])

		switch key {
		case "ProcessId":
			pid, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			processes = append(processes, newWindowsProcess(pid, ""))
		case "CommandLine":
			if len(processes) > 0 {
				processes[len(processes)-1].(*WindowsProcess).exe = value
			}
		}
	}

	return processes, nil
}
