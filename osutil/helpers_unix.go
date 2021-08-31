// +build linux

/*
 * Author: fasion
 * Created time: 2019-05-27 08:10:08
 * Last Modified by: fasion
 * Last Modified time: 2021-08-31 09:45:39
 */

package osutil

import (
	"os"
	"syscall"
)

func DupTo(path string, mode int, perm uint32, fd int) error {
	newfd, err := syscall.Open(path, mode, perm)
	if err != nil {
		return err
	}
	defer syscall.Close(newfd)

	return syscall.Dup2(newfd, fd)
}

func LockFile(f *os.File, how int) error {
	return syscall.Flock(int(f.Fd()), how)
}

func LockFileExclusive(f *os.File) error {
	return LockFile(f, syscall.LOCK_EX)
}

func LockFileShared(f *os.File) error {
	return LockFile(f, syscall.LOCK_SH)
}

func TryLockFile(f *os.File, how int) error {
	return LockFile(f, how|syscall.LOCK_NB)
}
