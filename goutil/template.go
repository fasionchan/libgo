/*
 * Author: fasion
 * Created time: 2021-10-12 19:21:32
 * Last Modified by: fasion
 * Last Modified time: 2021-11-19 16:19:30
 */

package goutil

import (
	"reflect"
	"text/template"
	"time"
)

type TemplateFuncMap template.FuncMap

func (m TemplateFuncMap) Native() template.FuncMap {
	return template.FuncMap(m)
}

func (m TemplateFuncMap) Dup() TemplateFuncMap {
	result := TemplateFuncMap{}
	for k, v := range m {
		result[k] = v
	}
	return result
}

func (m TemplateFuncMap) With(key string, f interface{}) TemplateFuncMap {
	m[key] = f
	return m
}

var TemplateHelpers = TemplateFuncMap{
	"limit": func(data interface{}, n int) interface{} {
		if n < 0 {
			n = 0
		}

		value := reflect.ValueOf(data)
		if value.Len() <= n {
			return data
		}

		return value.Slice(0, n).Interface()
	},
	"timefmt": func(t time.Time, fmt string, zeroPlaceHolder string) string {
		if t.IsZero() {
			return zeroPlaceHolder
		}

		if fmt == "" {
			fmt = DefaultTimeFormat
		}

		return t.Format(fmt)
	},
}
