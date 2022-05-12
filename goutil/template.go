/*
 * Author: fasion
 * Created time: 2021-10-12 19:21:32
 * Last Modified by: fasion
 * Last Modified time: 2022-05-10 16:13:06
 */

package goutil

import (
	"encoding/json"
	"math"
	"net/http"
	"reflect"
	"strings"
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

func (m TemplateFuncMap) WithDataContainer(name string) *DataContainer {
	container := NewDataContainer()
	m["set"+name] = container.Set
	return container
}

func (m TemplateFuncMap) WithDataContainers(names ...string) DataContainers {
	containers := make(DataContainers, 0, len(names))
	for _, name := range names {
		containers = containers.Append(m.WithDataContainer(name))
	}
	return containers
}

func (m TemplateFuncMap) WithDataContainersForMapping(names ...string) DataContainerMappingByString {
	containers := m.WithDataContainers(names...)
	mapping := DataContainerMappingByString{}
	for i, name := range names {
		mapping[name] = containers[i]
	}
	return mapping
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
	"now":          time.Now,
	"today":        Today,
	"dayOf":        DayOf,
	"duration":     time.ParseDuration,
	"intraDayTime": ParseIntraDayTime,
	"timefmt": func(t time.Time, fmt string, zeroPlaceHolder string) string {
		if t.IsZero() {
			return zeroPlaceHolder
		}

		if fmt == "" {
			fmt = DefaultTimeFormat
		}

		return t.Format(fmt)
	},
	"jsonify": func(v interface{}) (string, error) {
		data, err := json.Marshal(v)
		return string(data), err
	},
	"jsondecode": func(j string) (interface{}, error) {
		var data interface{}
		if err := json.Unmarshal([]byte(j), &data); err != nil {
			return nil, err
		}
		return data, nil
	},
	"makeMap": func(k interface{}, v interface{}) interface{} {
		return reflect.MakeMap(reflect.MapOf(reflect.TypeOf(k), reflect.TypeOf(v))).Interface()
	},
	"setMap": func(m interface{}, k interface{}, v interface{}) interface{} {
		reflect.ValueOf(m).SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
		return m
	},
	"httpGetJson": func(url string) (interface{}, error) {
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		var data interface{}
		if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
			return nil, err
		}

		return data, nil
	},

	"convert": Convert,
	"int": func(v interface{}) interface{} {
		return Convert(v, "int")
	},
	"int64": func(v interface{}) interface{} {
		return Convert(v, "int64")
	},
	"float64": func(v interface{}) interface{} {
		return Convert(v, "float64")
	},

	"addInt":     AddToInt,
	"addInt64":   AddToInt64,
	"addFloat64": AddToFloat64,

	"hasSubstr": strings.Contains,
	"hasPrefix": strings.HasPrefix,
	"hasSuffix": strings.HasSuffix,
	"join":      strings.Join,
	"split": func(s string, sep string) StringSlice {
		return strings.Split(s, sep)
	},
	"toLower":       strings.ToLower,
	"toUpper":       strings.ToUpper,
	"replaceSubstr": strings.ReplaceAll,
	"trim":          strings.Trim,
	"trimSpace":     strings.TrimSpace,

	"abs": math.Abs,
	"toStringSlice": func(s []string) StringSlice {
		return s
	},
}

type DataContainer struct {
	d interface{}
}

func NewDataContainer() *DataContainer {
	return &DataContainer{}
}

func (c *DataContainer) Get() interface{} {
	return c.d
}

func (c *DataContainer) Set(d interface{}) *DataContainer {
	c.d = d
	return c
}

type DataContainers []*DataContainer

func (containers DataContainers) Append(others ...*DataContainer) DataContainers {
	return append(containers, others...)
}

type DataContainerMappingByString map[string]*DataContainer
