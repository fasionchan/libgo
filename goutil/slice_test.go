/*
 * Author: fasion
 * Created time: 2021-05-14 09:27:21
 * Last Modified by: fasion
 * Last Modified time: 2022-03-10 14:58:38
 */

package goutil

import (
	"fmt"
	"testing"
)

func TestDataSlice(t *testing.T) {
	src := DataSlice{1, 2, 3}
	var nums []int
	pnums := &nums
	src.ConvertTo(&pnums)
	fmt.Println(nums)
}

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
