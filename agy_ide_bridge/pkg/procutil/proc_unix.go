//go:build !windows

package procutil

import (
	"os/exec"
	"syscall"
	"time"
)

// SetProcessGroup sets the process group for posix systems.
func SetProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

// KillProcessTree kills the process and its process group cleanly on posix systems.
func KillProcessTree(cmd *exec.Cmd, exited <-chan struct{}) {
	if cmd == nil || cmd.Process == nil {
		return
	}
	pid := cmd.Process.Pid
	pgid, err := syscall.Getpgid(pid)
	if err == nil {
		_ = syscall.Kill(-pgid, syscall.SIGTERM)
	} else {
		_ = cmd.Process.Signal(syscall.SIGTERM)
	}

	if exited != nil {
		select {
		case <-exited:
			return
		case <-time.After(1500 * time.Millisecond):
		}
	}

	if pgid > 0 {
		_ = syscall.Kill(-pgid, syscall.SIGKILL)
	}
	_ = cmd.Process.Kill()
}
