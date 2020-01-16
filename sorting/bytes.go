/*
 * Author: fasion
 * Created time: 2019-05-21 17:08:49
 * Last Modified by: fasion
 * Last Modified time: 2019-05-21 17:18:32
 */

package sorting

import (
	"bytes"
	"sort"
)

type sortByteSlices [][]byte

func (s sortByteSlices) Len() int {
	return len(s)
}

func (s sortByteSlices) Less(i, j int) bool {
	if bytes.Compare(s[i], s[j]) == -1 {
		return true
	}
	return false
}

func (s sortByteSlices) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func ByteSlices(raw [][]byte) ([][]byte) {
	sorted := sortByteSlices(raw)
	sort.Sort(sorted)
	return sorted
}