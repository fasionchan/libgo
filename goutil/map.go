/*
 * Author: fasion
 * Created time: 2022-03-10 13:59:02
 * Last Modified by: fasion
 * Last Modified time: 2022-03-10 15:17:01
 */

package goutil

import (
	"strings"
)

type DataMappingByPath map[string]interface{}

func (m DataMappingByPath) GetByPrefix(prefix string) (result DataSlice) {
	for path, data := range m {
		if strings.HasPrefix(path, prefix) {
			result = append(result, data)
		}
	}
	return
}

type DatasMappingByPath map[string]DataSlice

func (m DatasMappingByPath) GetByPrefix(prefix string) (result DataSlice) {
	for path, datas := range m {
		if strings.HasPrefix(path, prefix) {
			result = append(result, datas...)
		}
	}
	return
}
