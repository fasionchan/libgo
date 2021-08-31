/*
 * Author: fasion
 * Created time: 2019-05-27 09:46:53
 * Last Modified by: fasion
 * Last Modified time: 2021-08-31 09:45:37
 */

package osutil

import (
	"fmt"
	"os"
	"strings"
)

type EnvironMap map[string]string

func (self EnvironMap) Copy() EnvironMap {
	copy := EnvironMap{}

	for key, value := range self {
		copy[key] = value
	}

	return copy
}

func (self EnvironMap) With(name, value string) EnvironMap {
	self[name] = value
	return self
}

func (self EnvironMap) WithAll(all EnvironMap) EnvironMap {
	for name, value := range all {
		self[name] = value
	}
	return self
}

func (self EnvironMap) WithCurrent(names ...string) EnvironMap {
	if len(names) == 0 {
		for _, item := range os.Environ() {
			fields := strings.Split(item, "=")
			self[fields[0]] = fields[1]
		}
	} else {
		for _, name := range names {
			if value, ok := os.LookupEnv(name); ok {
				self[name] = value
			}
		}
	}

	return self
}

func (self EnvironMap) Format() []string {
	environ := make([]string, 0, len(self))
	for name, value := range self {
		environ = append(environ, fmt.Sprintf("%s=%s", name, value))
	}

	return environ
}
