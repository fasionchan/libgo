// +build linux

/*
 * Author: fasion
 * Created time: 2019-05-26 20:37:05
 * Last Modified by: fasion
 * Last Modified time: 2019-12-23 10:24:46
 */

package os

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var _ = fmt.Println

const (
	STATE_VAR_NAME = "_DAEMONIZE_STATE"
	BEFORE_FORK    = ""
	FIRST_FORK     = "FIRST_FORK"
	SECOND_FORK    = "SECOND_FORK"
)

func BatchCloseOnExec(start, end int) {
	for start < end {
		syscall.CloseOnExec(start)
		start += 1
	}
}

func beforeFork(envm EnvironMap) error {
	attr := syscall.ProcAttr{
		Env: envm.With(STATE_VAR_NAME, FIRST_FORK).Format(),
		Files: []uintptr{
			uintptr(syscall.Stdin),
			uintptr(syscall.Stdout),
			uintptr(syscall.Stderr),
		},
	}

	BatchCloseOnExec(3, 1024)
	_, err := syscall.ForkExec(os.Args[0], os.Args, &attr)
	if err != nil {
		return err
	}

	os.Exit(0)

	return nil
}

func firstFork() error {
	_, err := syscall.Setsid()
	if err != nil {
		return err
	}

	attr := syscall.ProcAttr{
		Env: EnvironMap{}.WithCurrent().With(STATE_VAR_NAME, SECOND_FORK).Format(),
		Files: []uintptr{
			uintptr(syscall.Stdin),
			uintptr(syscall.Stdout),
			uintptr(syscall.Stderr),
		},
	}

	BatchCloseOnExec(3, 1024)
	_, err = syscall.ForkExec(os.Args[0], os.Args, &attr)
	if err != nil {
		return err
	}

	os.Exit(0)

	return nil
}

func secondFork(stdin, stdout, stderr string) error {
	// change current working directory to root
	err := syscall.Chdir("/")
	if err != nil {
		return err
	}

	// clear file creation mask
	syscall.Umask(0)

	// ignore signal HUP
	signal.Ignore(syscall.SIGHUP)

	if false {
		limit := syscall.Rlimit{}
		err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit)
		if err != nil {
			return err
		}

		maxfds := limit.Max
		/*
			if maxfds == syscall.RLIM_INFINITY {
				maxfds = 1024
			}
		*/

		// close all open files
		for fd := maxfds - maxfds; fd < maxfds; fd++ {
			switch int(fd) {
			case syscall.Stdin:
			case syscall.Stdout:
			case syscall.Stderr:
			default:
				syscall.Close(int(fd))
			}
		}
	}

	// reopen stdin, stdout, and stderr
	err = ReopenStdio(stdin, stdout, stderr)
	if err != nil {
		return err
	}

	return nil
}

func Daemonize(stdin, stdout, stderr string, envm EnvironMap) error {
	switch GetDaemonizationState() {
	case BEFORE_FORK:
		return beforeFork(envm)
	case FIRST_FORK:
		return firstFork()
	case SECOND_FORK:
		return secondFork(stdin, stdout, stderr)
	}

	return nil
}

func GetDaemonizationState() string {
	return os.Getenv(STATE_VAR_NAME)
}

func DaemonForked() bool {
	switch GetDaemonizationState() {
	case BEFORE_FORK:
		return false
	case FIRST_FORK:
		return true
	case SECOND_FORK:
		return false
	}

	return false
}
