/*
 * Author: fasion
 * Created time: 2021-10-12 19:21:32
 * Last Modified by: fasion
 * Last Modified time: 2022-03-07 14:35:51
 */

package goutil

import (
	"encoding/json"
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
	"now":      time.Now,
	"today":    Today,
	"dayOf":    DayOf,
	"duration": time.ParseDuration,
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
