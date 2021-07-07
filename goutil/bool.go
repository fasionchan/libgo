/*
 * Author: fasion
 * Created time: 2021-07-07 08:19:00
 * Last Modified by: fasion
 * Last Modified time: 2021-07-07 13:38:59
 */

package goutil

type BoolMappingByString map[string]bool

func NewBoolMappingByString() BoolMappingByString {
	return BoolMappingByString{}
}

func (mapping BoolMappingByString) Dup() BoolMappingByString {
	return NewBoolMappingByString().Merge(mapping)
}

func (mapping BoolMappingByString) FilterByKeys(keys ...string) BoolMappingByString {
	result := BoolMappingByString{}
	for _, key := range keys {
		if value, ok := mapping[key]; ok {
			result[key] = value
		}
	}
	return result
}

func (mapping BoolMappingByString) FilterByKeysWithDefault(defaultValue bool, keys ...string) BoolMappingByString {
	return mapping.FilterByKeys(keys...).SetDefault(defaultValue, keys...)
}

func (mapping BoolMappingByString) Merge(other BoolMappingByString) BoolMappingByString {
	for key, value := range other {
		mapping[key] = value
	}
	return mapping
}

func (mapping BoolMappingByString) SetDefault(defaultValue bool, keys ...string) BoolMappingByString {
	for _, key := range keys {
		if ok, _ := mapping[key]; !ok {
			mapping[key] = defaultValue
		}
	}
	return mapping
}

func (mapping BoolMappingByString) BatchSet(value bool, keys ...string) BoolMappingByString {
	for _, key := range keys {
		mapping[key] = value
	}
	return mapping
}

func (mapping BoolMappingByString) AnyTrue() bool {
	for _, value := range mapping {
		if value {
			return true
		}
	}

	return false
}

func (mapping BoolMappingByString) AnyKeysTrue(keys ...string) bool {
	for _, key := range keys {
		if value := mapping[key]; value {
			return true
		}
	}

	return false
}
