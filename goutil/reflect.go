/*
 * Author: fasion
 * Created time: 2022-03-08 11:40:21
 * Last Modified by: fasion
 * Last Modified time: 2022-03-08 12:04:25
 */

package goutil

import "reflect"

var (
	IntType     = reflect.TypeOf(int(0))
	Int64Type   = reflect.TypeOf(int64(0))
	Float64Type = reflect.TypeOf(float64(0))
)

var BasicTypeMapping = map[string]reflect.Type{
	"int":   reflect.TypeOf(int(0)),
	"int8":  reflect.TypeOf(int8(0)),
	"int16": reflect.TypeOf(int16(0)),
	"int32": reflect.TypeOf(int32(0)),
	"int64": reflect.TypeOf(int64(0)),

	"uint":   reflect.TypeOf(uint(0)),
	"uint8":  reflect.TypeOf(uint8(0)),
	"uint16": reflect.TypeOf(uint16(0)),
	"uint32": reflect.TypeOf(uint32(0)),
	"uint64": reflect.TypeOf(uint64(0)),

	"float32": reflect.TypeOf(float32(0)),
	"float64": reflect.TypeOf(float64(0)),
}

func Convert(v interface{}, t string) interface{} {
	return reflect.ValueOf(v).Convert(BasicTypeMapping[t]).Interface()
}

func AddToInt(v int, others ...interface{}) int {
	for _, other := range others {
		v += int(reflect.ValueOf(other).Convert(IntType).Int())
	}
	return v
}

func AddToInt64(v int64, others ...interface{}) int64 {
	for _, other := range others {
		v += reflect.ValueOf(other).Convert(Int64Type).Int()
	}
	return v
}

func AddToFloat64(v float64, others ...interface{}) float64 {
	for _, other := range others {
		v += reflect.ValueOf(other).Convert(Float64Type).Float()
	}
	return v
}
