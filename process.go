/*

Original Repository
https://github.com/mitchellh/go-ps

This version is just for personal usages.. "wink wink"
So, No support for this

*/

package process

import (
	"regexp"
	"strings"
)

type Process interface {
	// Pid is the process ID for this process.
	Pid() int
	// Executable name with the entire command.
	Executable() string
}

func Processes() ([]Process, error) {
	return processes()
}

// Global BlackList
var blackList = []string{
	"malware69",
}

// Set new BlackList
func SetBlackList(newBlackList []string) {
	blackList = newBlackList
}

func detectFunc(blacklist, list []string) string {
	for _, blacklisted := range blacklist {
		curr := regexp.MustCompile("(?i)" + blacklisted)
		for _, process := range list {
			if curr.MatchString(process) {
				return blacklisted
			}
		}
	}
	return ""
}

// Compare running process with BlackList
func CheckProcess() string {
	processList, _ := Processes()
	list := make([]string, len(processList))
	for i, p := range processList {
		list[i] = strings.ToLower(p.Executable())
	}
	return detectFunc(blackList, list)
}
