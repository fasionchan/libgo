/*
 * Author: fasion
 * Created time: 2021-07-01 18:50:08
 * Last Modified by: fasion
 * Last Modified time: 2021-11-03 14:42:41
 */

package goutil

import (
	"bytes"
	"fmt"
	"reflect"
)

type dataNode struct {
	Key      string `json:"Key"`
	TypeName string `json:"Type"`

	ptr  reflect.Value
	v    reflect.Value
	t    reflect.Type
	data interface{}
}

func newDataNode(key, typeName string, ptr reflect.Value, data interface{}) *dataNode {
	return &dataNode{
		Key:      key,
		TypeName: typeName,
		ptr:      ptr,
		data:     data,
	}
}

func dataNodeFromValue(key string, v reflect.Value) (*dataNode, error) {
	var ptr reflect.Value

	data := v.Interface()
	t := v.Type()
	for t.Kind() == reflect.Interface {
		v = v.Elem()
		t = v.Type()
	}

	if t.Kind() == reflect.Ptr {
		ptr = v
	} else {
		// data should be addressable
		if !v.CanAddr() {
			// fmt.Println("not addressable", key)
			// fmt.Println("not addressable", v)
			// fmt.Println("not addressable", v.Interface())
			return nil, fmt.Errorf("data should be addressable")
		}

		ptr = v.Addr()
	}

	// fmt.Println("asdf", ptr)

	typeName := t.String()

	return newDataNode(key, typeName, ptr, data).derefer(), nil
}

func dataNodeFromInterface(key string, data interface{}) (*dataNode, error) {
	return dataNodeFromValue(key, reflect.ValueOf(data))
}

func (node *dataNode) Ident() string {
	return fmt.Sprintf("%d(%s)", node.ptr.Pointer(), node.t.String())
}

func (node *dataNode) derefer() *dataNode {
	for !node.ptr.IsNil() {
		node.v = node.ptr.Elem()
		node.t = node.v.Type()

		if node.t.Kind() != reflect.Ptr {
			return node
		}

		node.ptr = node.v
	}

	return nil
}

func (node *dataNode) ReadableString() string {
	return fmt.Sprintf("%s(%s)", node.Key, node.TypeName)
}

type dataNodes []*dataNode

func (nodes dataNodes) Len() int {
	return len(nodes)
}

func (nodes dataNodes) Empty() bool {
	return nodes.Len() == 0
}

func (nodes dataNodes) Last() *dataNode {
	if nodes.Empty() {
		return nil
	}

	return nodes[nodes.Len()-1]
}

func (nodes dataNodes) Dup() dataNodes {
	result := make(dataNodes, 0, len(nodes))
	for _, node := range nodes {
		result = result.Append(node)
	}
	return result
}

func (nodes dataNodes) Append(more ...*dataNode) dataNodes {
	return append(nodes, more...)
}

func (nodes dataNodes) ReadableString(highlight int) string {
	buf := bytes.NewBuffer(nil)

	for i, node := range nodes {
		if i > 0 {
			fmt.Fprintf(buf, " => ")
		}

		if i == highlight {
			fmt.Fprint(buf, "@")
		}

		fmt.Fprint(buf, node.ReadableString())
	}

	return buf.String()
}

type cyclePath struct {
	Path     dataNodes `json:"Path"`
	Entrance int       `json:"Entrance"`
}

func (path cyclePath) ReadableString() string {
	return path.Path.ReadableString(path.Entrance)
}

type cyclePathes []*cyclePath

func (pathes cyclePathes) Print() {
	fmt.Println(pathes.ReadableStrings())
}

func (pathes cyclePathes) Len() int {
	return len(pathes)
}

func (pathes cyclePathes) Limit(n int) cyclePathes {
	if pathes.Len() <= n {
		return pathes
	}

	return pathes[:n]
}

func (pathes cyclePathes) Empty() bool {
	return pathes.Len() == 0
}

func (pathes cyclePathes) ReadableStrings() StringSlice {
	parts := make(StringSlice, 0, pathes.Len())
	for _, path := range pathes {
		parts = parts.Append(path.ReadableString())
	}
	return parts
}

func (pathes cyclePathes) Append(more ...*cyclePath) cyclePathes {
	return append(pathes, more...)
}

type cycleDetector struct {
	path         dataNodes
	layerMapping map[interface{}]int
	current      *dataNode

	cyclePathes cyclePathes
}

func newCycleDetector() *cycleDetector {
	return &cycleDetector{
		layerMapping: map[interface{}]int{},
	}
}

func (detector *cycleDetector) push(node *dataNode) {
	detector.layerMapping[node.Ident()] = detector.path.Len()
	detector.path = detector.path.Append(node)
	detector.current = node
}

func (detector *cycleDetector) pushIfNoCycle(node *dataNode) bool {
	// fmt.Println("push")
	// fmt.Println("mapping", detector.layerMapping)
	// fmt.Println("node", node)
	// fmt.Println("Ptr", node.ptr)
	// fmt.Println("t", node.t.String())
	// fmt.Println("data", node.data)

	if i, ok := detector.layerMapping[node.Ident()]; ok {
		detector.cyclePathes = detector.cyclePathes.Append(&cyclePath{
			Path:     detector.path.Dup().Append(node),
			Entrance: i,
		})

		return false
	}

	detector.push(node)

	return true
}

func (detector *cycleDetector) pop() *dataNode {
	last := detector.path.Last()
	delete(detector.layerMapping, last.Ident())
	detector.path = detector.path[:detector.path.Len()-1]
	detector.current = detector.path.Last()
	return last
}

func (detector *cycleDetector) traverse(node *dataNode) error {
	// fmt.Println("traverse", detector.path.ReadableString(-1), ">>>>", node.ReadableString())
	// fmt.Println("key", node.Key)
	// fmt.Println("type", node.TypeName)
	// fmt.Println("data", node.data)
	// fmt.Println("ptr", node.ptr)
	// fmt.Println("v", node.v)
	// fmt.Println("t", node.t)
	// fmt.Println()

	// switch node.t.K

	if ok := detector.pushIfNoCycle(node); !ok {
		return nil
	}
	defer detector.pop()

	v := node.v
	t := node.t

	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			// fmt.Println("field", field.Name)

			// ignore unexported
			if field.PkgPath != "" {
				continue
			}

			// fmt.Println("field")
			// fmt.Println("field", v.Field(i))
			// fmt.Println("field", v.Field(i).Type())
			// fmt.Println("field", v.Field(i).Type().Kind())

			node, err := dataNodeFromValue(field.Name, v.Field(i))
			if err != nil {
				return err
			}

			if node != nil {
				detector.traverse(node)
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			node, err := dataNodeFromValue(fmt.Sprintf("%d", i), v.Index(i))
			if err != nil {
				return err
			}

			if node != nil {
				detector.traverse(node)
			}
		}
	case reflect.Map:
	}

	return nil
}

func (detector *cycleDetector) Detect(data interface{}) error {
	node, err := dataNodeFromInterface("", data)
	if err != nil {
		return err
	}

	return detector.traverse(node)
	return nil
}

func DetectCycle(v interface{}) (result cyclePathes, err error) {
	detector := newCycleDetector()
	err = detector.Detect(v)
	result = detector.cyclePathes
	return
}

func MustDetectCycle(v interface{}) cyclePathes {
	result, err := DetectCycle(v)
	if err != nil {
		panic(err)
	}
	return result
}
