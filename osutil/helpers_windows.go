/*
 * Author: fasion
 * Created time: 2019-08-15 16:25:59
 * Last Modified by: fasion
 * Last Modified time: 2021-08-31 09:45:42
 */

package osutil

import (
	"syscall"
)

var (
	kernel32 = syscall.MustLoadDLL("Kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

func DupTo(path string, mode int, perm uint32, fd syscall.Handle) (error) {
	newfd, err := syscall.Open(path, mode, perm)
	if err != nil {
		return err
	}
	defer syscall.Close(newfd)

	return nil
	// return syscall.Dup2(newfd, fd)
}
