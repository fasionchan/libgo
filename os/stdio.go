/*
 * Author: fasion
 * Created time: 2019-08-15 16:34:51
 * Last Modified by: fasion
 * Last Modified time: 2019-08-15 16:37:25
 */

package os

import (
	"syscall"
)

func ReopenStdio(stdin, stdout, stderr string) (error) {
	if stdin != "" {
		if err := DupTo(stdin, syscall.O_RDONLY, 0, syscall.Stdin); err != nil {
			return err
		}
	}

	if stdout != "" {
		if err := DupTo(stdout, syscall.O_WRONLY | syscall.O_APPEND, 0, syscall.Stdout); err != nil {
			return err
		}
	}

	if stderr != "" {
		if err := DupTo(stderr, syscall.O_WRONLY | syscall.O_APPEND, 0, syscall.Stderr); err != nil {
			return err
		}
	}

	return nil
}
