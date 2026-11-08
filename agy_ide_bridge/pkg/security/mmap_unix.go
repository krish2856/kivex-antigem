//go:build !windows

package security

import "syscall"

func mmapFile(fd int, size int) ([]byte, error) {
	return syscall.Mmap(fd, 0, size, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
}

func munmapFile(data []byte) error {
	return syscall.Munmap(data)
}
