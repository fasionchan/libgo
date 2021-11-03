/*
 * Author: fasion
 * Created time: 2021-05-14 09:27:21
 * Last Modified by: fasion
 * Last Modified time: 2021-11-02 18:53:29
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

func TestReverseSlice(t *testing.T) {
	var case1 = []int{0, 1, 2, 3}
	ReverseSlice(case1)
	fmt.Println(case1)

	var case2 = &[]int{0, 1, 2, 3, 4}
	ReverseSlice(&case2)
	fmt.Println(case2)
}
