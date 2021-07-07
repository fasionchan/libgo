/*
 * Author: fasion
 * Created time: 2021-02-09 10:16:51
 * Last Modified by: fasion
 * Last Modified time: 2021-07-07 08:55:10
 */

package goutil

import (
	"sort"
	"strings"
)

type StringSlice []string

func NewStringSlice(ss ...string) StringSlice {
	return StringSlice(ss)
}

func NewStringSliceFromStrings(ss []string) StringSlice {
	return StringSlice(ss)
}

func (ss StringSlice) TrueMapping() BoolMappingByString {
	return NewBoolMappingByString().BatchSet(true, ss...)
}

func (ss StringSlice) FalseMapping() BoolMappingByString {
	return NewBoolMappingByString().BatchSet(false, ss...)
}

func (ss StringSlice) ToSet() StringSet {
	return NewStringSet(ss...)
}

func (ss StringSlice) Strings() []string {
	return []string(ss)
	// result := make([]string, 0, len(ss))
	// for _, s := range ss {
	// 	result = append(result, s)
	// }
	// return result
}

func (ss StringSlice) Map(f func(string) string) StringSlice {
	for i, s := range ss {
		ss[i] = f(s)
	}
	return ss
}

func (ss StringSlice) Filter(f func(string) bool) StringSlice {
	result := make(StringSlice, 0, len(ss))
	for _, s := range ss {
		if f(s) {
			result = result.Append(s)
		}
	}
	return result
}

func (ss StringSlice) FilterNot(f func(string) bool) StringSlice {
	return ss.Filter(func(s string) bool {
		return !f(s)
	})
}

func (ss StringSlice) PurgeZero() StringSlice {
	return ss.Filter(func(s string) bool {
		return s != ""
	})
}

func (ss StringSlice) split(sep string) StringSlice {
	result := make(StringSlice, 0, len(ss))

	for _, s := range ss {
		for _, sub := range strings.Split(s, sep) {
			result = result.Append(sub)
		}
	}

	return result
}

func (ss StringSlice) Split(seps ...string) StringSlice {
	for _, sep := range seps {
		ss = ss.split(sep)
	}
	return ss
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

func (ss StringSlice) Empty() bool {
	return ss.Len() == 0
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

type StringSet map[string]bool

func NewStringSet(ss ...string) StringSet {
	set := StringSet{}
	for _, s := range ss {
		set[s] = true
	}
	return set
}

func (set StringSet) Dup() StringSet {
	result := StringSet{}
	for s := range set {
		result[s] = true
	}
	return result
}

func (set StringSet) Len() int {
	return len(set)
}

func (set StringSet) Empty() bool {
	return set.Len() == 0
}

func (set StringSet) ToSlice() StringSlice {
	slice := make(StringSlice, 0, len(set))
	for s := range set {
		slice = slice.Append(s)
	}
	return slice
}

func (a StringSet) Equal(b StringSet) bool {
	if len(a) != len(b) {
		return false
	}

	for s := range a {
		if !b.Contain(s) {
			return false
		}
	}

	return true
}

func (set StringSet) Contain(s string) (ok bool) {
	_, ok = set[s]
	return
}

func (set StringSet) AddStrings(ss ...string) StringSet {
	for _, s := range ss {
		set[s] = true
	}
	return set
}

func (a StringSet) Add(b StringSet) StringSet {
	for s := range b {
		a[s] = true
	}
	return a
}

func (a StringSet) Sub(b StringSet) StringSet {
	result := StringSet{}
	for s := range a {
		if !b.Contain(s) {
			result[s] = true
		}
	}
	return result
}

func (a StringSet) Union(b StringSet) StringSet {
	result := a.Dup()
	for s := range b {
		result.AddStrings(s)
	}
	return result
}

func (a StringSet) Difference(b StringSet) StringSet {
	return a.Dup().Sub(b)
}

func (a StringSet) SymmetricDifference(b StringSet) StringSet {
	result := StringSet{}

	for s := range a {
		if !b.Contain(s) {
			result[s] = true
		}
	}

	for s := range b {
		if !a.Contain(s) {
			result[s] = true
		}
	}

	return result
}

func (a StringSet) Intersection(b StringSet) StringSet {
	result := StringSet{}
	for s := range a {
		if b.Contain(s) {
			result[s] = true
		}
	}
	return result
}

type StringSets []StringSet

type CounterByStringKey map[string]int

func NewCounterByStringKey() CounterByStringKey {
	return CounterByStringKey{}
}

func (counter CounterByStringKey) Increase(key string) {
	counter[key]++
}

func (counter CounterByStringKey) SumKeys(keys ...string) (total int) {
	for _, key := range keys {
		total += counter[key]
	}
	return
}
