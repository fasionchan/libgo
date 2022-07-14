/*
 * Author: fasion
 * Created time: 2022-07-14 14:27:50
 * Last Modified by: fasion
 * Last Modified time: 2022-07-14 14:48:47
 */

package generics

import (
	"testing"
)

type CompareOrderedSliceTestCase struct {
	a      IntSlice
	b      IntSlice
	result int
}

func TestCompareOrderedSlice(t *testing.T) {
	testCases := []struct {
		a      IntSlice
		b      IntSlice
		result int
	}{
		{
			a:      IntSlice{1, 2, 3},
			b:      IntSlice{1, 2, 3},
			result: 0,
		},
		{
			a:      IntSlice{1, 2},
			b:      IntSlice{1, 2, 3},
			result: -1,
		},
		{
			a:      IntSlice{},
			b:      IntSlice{1},
			result: -1,
		},
		{
			a:      IntSlice{},
			b:      IntSlice{},
			result: 0,
		},
		{
			a:      IntSlice{0, 0, 3},
			b:      IntSlice{0, 1},
			result: -1,
		},
	}

	for i, testCase := range testCases {
		a := testCase.a
		b := testCase.b
		expected := testCase.result

		result := a.Compare(b)
		if result != expected {
			t.Errorf("TestCompareOrderedSlice step ① failed: [%d] %v", i, testCase)
		}

		var less, greater, equal bool
		switch result {
		case -1:
			less = true
		case 0:
			equal = true
		case 1:
			greater = true
		}

		if a.Less(b) != less {
			t.Errorf("TestCompareOrderedSlice step ② failed: [%d] %v", i, testCase)
		}
		if a.Greater(b) != greater {
			t.Errorf("TestCompareOrderedSlice step ③ failed: [%d] %v", i, testCase)
		}
		if a.Equal(b) != equal {
			t.Errorf("TestCompareOrderedSlice step ④ failed: [%d] %v", i, testCase)
		}

		if result != 0 {
			if b.Compare(a) != -expected {
				t.Errorf("TestCompareOrderedSlice step ⑤ failed: [%d] %v", i, testCase)
			}
		}
	}
}
