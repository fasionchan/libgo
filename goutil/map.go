/*
 * Author: fasion
 * Created time: 2022-03-10 13:59:02
 * Last Modified by: fasion
 * Last Modified time: 2022-03-11 08:46:58
 */

package goutil

import (
	"strings"
)

type DataMappingByKey map[string]interface{}

func (m DataMappingByKey) Len() int {
	return len(m)
}

func (m DataMappingByKey) Empty() bool {
	return m.Len() == 0
}

func (m DataMappingByKey) Keys() StringSlice {
	keys := make(StringSlice, 0, len(m))
	for key, _ := range m {
		keys = keys.Append(key)
	}
	return keys
}

func (m DataMappingByKey) Values() DataSlice {
	values := make(DataSlice, 0, len(m))
	for _, value := range m {
		values = values.Append(value)
	}
	return values
}

func (m DataMappingByKey) Datas() DataSlice {
	return m.Values()
}

func (m DataMappingByKey) FilterByAnyPrefix(prefixes StringSlice) DataMappingByKey {
	result := DataMappingByKey{}
	for key, data := range m {
		if prefixes.AnyMatch(func(s string) bool {
			return strings.HasPrefix(key, s)
		}) {
			result[key] = data
		}
	}
	return result
}

func (m DataMappingByKey) FilterByPrefix(prefixes ...string) DataMappingByKey {
	return m.FilterByAnyPrefix(prefixes)
}

type DatasMappingByKey map[string]DataSlice

func (m DatasMappingByKey) Keys() StringSlice {
	keys := make(StringSlice, 0, len(m))
	for key, _ := range m {
		keys = keys.Append(key)
	}
	return keys
}

func (m DatasMappingByKey) FilterByAnyPrefix(prefixes StringSlice) DatasMappingByKey {
	result := DatasMappingByKey{}
	for key, datas := range m {
		if prefixes.AnyMatch(func(s string) bool {
			return strings.HasPrefix(key, s)
		}) {
			result[key] = datas
		}
	}
	return result
}
