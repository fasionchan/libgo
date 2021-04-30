/*
 * Author: fasion
 * Created time: 2021-04-26 13:32:28
 * Last Modified by: fasion
 * Last Modified time: 2021-04-26 13:49:27
 */

package goutil

import "testing"

type StringSetTestCase struct {
	a      StringSet
	b      StringSet
	result StringSet
}

type StringSetTestCases []*StringSetTestCase

func TestStringSetSub(t *testing.T) {
	testCases := StringSetTestCases{
		&StringSetTestCase{
			a:      NewStringSet("a", "b", "c"),
			b:      NewStringSet("b", "c", "d"),
			result: NewStringSet("a"),
		},
		&StringSetTestCase{
			a:      NewStringSet("a", "b", "c"),
			b:      NewStringSet(),
			result: NewStringSet("a", "b", "c"),
		},
		&StringSetTestCase{
			a:      NewStringSet(),
			b:      NewStringSet("b", "c", "d"),
			result: NewStringSet(),
		},
	}

	for i, testCase := range testCases {
		if !testCase.a.Dup().Sub(testCase.b).Equal(testCase.result) {
			t.Errorf("StringSet Sub test case %d failed", i)
		}
	}
}
