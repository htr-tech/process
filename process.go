/*

Original Repository
https://github.com/mitchellh/go-ps

This version is just for personal usages.. "wink wink"
So, No support for this

*/

package process

type Process interface {
	// Pid is the process ID for this process.
	Pid() int

	// Executable name with the entire command.
	Executable() string
}

func Processes() ([]Process, error) {
	return processes()
}
