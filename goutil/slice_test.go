/*
 * Author: fasion
 * Created time: 2021-05-14 09:27:21
 * Last Modified by: fasion
 * Last Modified time: 2021-05-14 09:39:32
 */

package goutil

import (
	"fmt"
	"testing"
)

func TestDupSlice(t *testing.T) {
	src := []string{"a", "b", "c"}

	var dst []string
	DupSlice(src, &dst)

	fmt.Println(src, dst)

	dst[0] = "A"
	fmt.Println(src, dst)

}
