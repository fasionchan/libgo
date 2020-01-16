/*
 * Author: fasion
 * Created time: 2019-12-09 16:18:51
 * Last Modified by: fasion
 * Last Modified time: 2019-12-09 16:19:38
 */

package arch

import (
	"encoding/binary"
)

func Htons(n uint16) uint16 {
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], n)
	return NativeEndian.Uint16(buf[:])
}

func Ntohs(n uint16) uint16 {
	var buf [2]byte
	NativeEndian.PutUint16(buf[:], n)
	return binary.BigEndian.Uint16(buf[:])
}

func Htonl(n uint32) uint32 {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], n)
	return NativeEndian.Uint32(buf[:])
}

func Ntohl(n uint32) uint32 {
	var buf [4]byte
	NativeEndian.PutUint32(buf[:], n)
	return binary.BigEndian.Uint32(buf[:])
}
