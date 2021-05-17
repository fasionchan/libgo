/*
 * Author: fasion
 * Created time: 2021-05-14 09:24:27
 * Last Modified by: fasion
 * Last Modified time: 2021-05-14 09:57:55
 */

package goutil

import (
	"reflect"
)

func DupSlice(src, dst interface{}) {
	// dst should be a pointer to slice
	dstPtr := reflect.ValueOf(dst)
	dstPtrType := dstPtr.Type()
	if dstPtrType.Kind() != reflect.Ptr {
		panic("dst must be a pointer to slice")
	}

	// dereference pointer to src slice
	srcSlice := reflect.ValueOf(src)
	for srcSlice.Type().Kind() == reflect.Ptr {
		srcSlice = srcSlice.Elem()
	}

	// src & dst slice should be same type
	srcSliceType := srcSlice.Type()
	if srcSliceType != dstPtrType.Elem() {
		panic("src and dst should be slices of the same type")
	}

	// copy elements one by one
	srcSliceLen := srcSlice.Len()
	dupSlice := reflect.MakeSlice(srcSliceType, srcSliceLen, srcSliceLen)
	for i := 0; i < srcSliceLen; i += 1 {
		dupSlice.Index(i).Set(srcSlice.Index(i))
	}

	// set pointer
	dstPtr.Elem().Set(dupSlice)
}
