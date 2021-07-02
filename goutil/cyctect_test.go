/*
 * Author: fasion
 * Created time: 2021-07-01 19:00:07
 * Last Modified by: fasion
 * Last Modified time: 2021-07-02 10:40:24
 */

package goutil

import (
	"fmt"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	a := &struct {
		B interface{}
	}{}

	b := &struct {
		A interface{}
	}{
		A: a,
	}

	a.B = b

	pathes, err := DetectCycle(a)
	fmt.Println(err)

	fmt.Println(pathes.ReadableStrings())
}
