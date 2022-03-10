/*
 * Author: fasion
 * Created time: 2021-05-14 09:24:27
 * Last Modified by: fasion
 * Last Modified time: 2022-03-10 14:57:19
 */

package goutil

import (
	"reflect"
)

type DataSlice []interface{}

func (datas DataSlice) Append(others ...interface{}) DataSlice {
	return append(datas, others...)
}

func (datas DataSlice) ConvertTo(dst interface{}) error {
	dstValue := reflect.ValueOf(dst)
	dstType := dstValue.Type()

	var ptrValue reflect.Value
	for dstType.Kind() == reflect.Ptr {
		ptrValue = dstValue
		dstValue = dstValue.Elem()
		dstType = dstValue.Type()
	}

	if dstType.Kind() != reflect.Slice {
		return nil
	}

	dataType := dstType.Elem()
	for i, data := range datas {
		dataValue := reflect.ValueOf(data)

		if i == 0 {
			if dataValue.Type() != dataType {
				return nil
			}
		}

		dstValue = reflect.Append(dstValue, dataValue)
	}

	if ptrValue.IsValid() {
		ptrValue.Elem().Set(dstValue)
	}

	return nil
}

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

func ReverseSlice(slice interface{}) {
	value := reflect.ValueOf(slice)
	for value.Type().Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Type().Kind() != reflect.Slice {
		panic("data given is not a slice")
	}

	n := value.Len()
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		tmp := value.Index(i).Interface()
		value.Index(i).Set(value.Index(j))
		value.Index(j).Set(reflect.ValueOf(tmp))
	}
}
