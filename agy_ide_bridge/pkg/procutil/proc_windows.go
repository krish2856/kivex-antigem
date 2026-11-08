//go:build windows

package procutil

import (
	"fmt"
	"os/exec"
)

// SetProcessGroup is a no-op on Windows.
func SetProcessGroup(cmd *exec.Cmd) {
	// No Setpgid on Windows
}

// KillProcessTree kills the process and its children on Windows.
func KillProcessTree(cmd *exec.Cmd, exited <-chan struct{}) {
	if cmd == nil || cmd.Process == nil {
		return
	}
	pid := cmd.Process.Pid
	_ = cmd.Process.Kill()
	_ = exec.Command("taskkill", "/F", "/T", "/PID", fmt.Sprintf("%d", pid)).Run()
}
