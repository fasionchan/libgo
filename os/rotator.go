// +build linux

/*
 * Author: fasion
 * Created time: 2019-05-27 13:18:54
 * Last Modified by: fasion
 * Last Modified time: 2019-08-19 15:09:00
 */

package os

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
	"compress/gzip"
)

type SizedFileRotator struct {
	Fd int
	Path string
	Size int64
	Backups int
	Duration time.Duration
	CheckDuration time.Duration
	mutex sync.Mutex
}

func NewSizedFileRotator(fd int, path string, size int64, backups int, duration time.Duration, checkDuration time.Duration) (*SizedFileRotator, error) {
	rotator := &SizedFileRotator{
		Fd: fd,
		Path: path,
		Size: size,
		Backups: backups,
		Duration: duration,
		CheckDuration: checkDuration,
	}

	if err := rotator.createDup(); err != nil {
		return nil, err
	}

	go rotator.RotateForever()

	return rotator, nil
}

func (self *SizedFileRotator) createDup() (error) {
	return DupTo(self.Path, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644, self.Fd)
}

func (self *SizedFileRotator) RotateIfNeeded() (error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	curTime := time.Now()

	info, err := os.Stat(self.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return self.createDup()
		}
		return err
	}

	if info.Size() < self.Size {
		return nil
	}

	dir := filepath.Dir(self.Path)
	base := filepath.Base(self.Path)

	for i := self.Backups; i > 0; i-- {
		curName := fmt.Sprintf("%s/%s.%d.gz", dir, base, i)
		nextName := fmt.Sprintf("%s/%s.%d.gz", dir, base, i+1)

		info, err := os.Stat(curName)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}

		if i == self.Backups || curTime.Sub(info.ModTime()) > self.Duration {
			err = os.Remove(curName)
			if err != nil {
				return err
			}
		} else {
			err = os.Rename(curName, nextName)
			if err != nil {
				return err
			}
		}
	}

	backupName := fmt.Sprintf("%s/%s.1", dir, base)
	err = os.Rename(self.Path, backupName)
	if err != nil {
		return err
	}

	err = self.createDup()
	if err != nil {
		return err
	}

	frFile, err := os.Open(backupName)
	if err != nil {
		return err
	}
	defer frFile.Close()

	toFile, err := os.OpenFile(backupName + ".gz", os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer toFile.Close()

	zipWriter := gzip.NewWriter(toFile)

	_, err = io.Copy(zipWriter, frFile)
	if err != nil {
		return err
	}

	zipWriter.Close()

	err = os.Remove(backupName)
	if err != nil {
		return err
	}

	return nil
}

func (self *SizedFileRotator) RotateForever() {
	for {
		self.RotateIfNeeded()
		time.Sleep(self.CheckDuration)
	}
}
