/*
 * Author: fasion
 * Created time: 2021-02-09 10:16:51
 * Last Modified by: fasion
 * Last Modified time: 2021-03-18 14:29:09
 */

package goutil

import (
	"sort"
	"strings"
)

type StringSlice []string

func (ss StringSlice) Strings() []string {
	result := make([]string, 0, len(ss))
	for _, s := range ss {
		result = append(result, s)
	}
	return result
}

func (ss StringSlice) Join(sep string) string {
	return strings.Join(ss.Strings(), sep)
}

func (ss StringSlice) Dup() StringSlice {
	dup := make(StringSlice, 0, len(ss))
	return append(dup, ss...)
}

func (ss StringSlice) Append(s ...string) StringSlice {
	return append(ss, s...)
}

func (ss StringSlice) Concat(others ...StringSlice) StringSlice {
	for _, other := range others {
		ss = ss.Append(other...)
	}
	return ss
}

func (ss StringSlice) Len() int {
	return len(ss)
}

func (ss StringSlice) Less(i, j int) bool {
	return ss[i] < ss[j]
}

func (ss StringSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss StringSlice) InplaceSort() StringSlice {
	sort.Sort(ss)
	return ss
}

func (ss StringSlice) Sort() StringSlice {
	return ss.Dup().InplaceSort()
}

func (ss StringSlice) Unique() StringSlice {
	result := make(StringSlice, 0, len(ss))

	var last string
	for i, s := range ss {
		if i == 0 || s != last {
			result = append(result, s)
			last = s
		}
	}

	return result
}

func (ss StringSlice) Equal(other StringSlice) bool {
	if len(ss) != len(other) {
		return false
	}

	if len(ss) == 0 {
		return true
	}

	other = other.Sort()
	for i, s := range ss.Sort() {
		if s != other[i] {
			return false
		}
	}

	return true
}

func (ss StringSlice) Contain(s string) bool {
	for _, _s := range ss {
		if _s == s {
			return true
		}
	}
	return false
}

func (ss StringSlice) Index(s string) int {
	for i, _s := range ss {
		if _s == s {
			return i
		}
	}

	return -1
}

func (ss StringSlice) Before(a, b string) bool {
	ai := ss.Index(a)
	bi := ss.Index(b)

	if ai != -1 && bi != -1 && ai < bi {
		return true
	}

	return false
}

func StringSliceToBooleanMapping(ss []string) map[string]bool {
	mapping := make(map[string]bool)
	for _, s := range ss {
		mapping[s] = true
	}
	return mapping
}

func StringSliceDifferenceSet(a, b []string) []string {
	mapping := StringSliceToBooleanMapping(b)
	result := make([]string, 0, len(a))
	for _, s := range a {
		if !mapping[s] {
			result = append(result, s)
		}
	}
	return result
}

func NewStringSlice(ss []string) StringSlice {
	return StringSlice(ss)
}
