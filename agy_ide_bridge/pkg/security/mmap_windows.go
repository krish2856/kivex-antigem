//go:build windows

package security

import "errors"

func mmapFile(fd int, size int) ([]byte, error) {
	return nil, errors.New("mmap not supported on windows")
}

func munmapFile(data []byte) error {
	return nil
}
