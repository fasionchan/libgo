/*
 * Author: fasion
 * Created time: 2019-08-06 16:21:24
 * Last Modified by: fasion
 * Last Modified time: 2019-08-07 08:29:32
 */

package encoding

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var _ = fmt.Println

func MarshalBinary(order binary.ByteOrder, datas... interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	for _, data := range datas {
		if err := binary.Write(buf, order, data); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func UnmarshalBinary(order binary.ByteOrder, data []byte, dst interface{}) (error) {
	r := bytes.NewReader(data)
	return binary.Read(r, order, dst)
}
