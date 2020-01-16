/**
 * FileName:   byte.go
 * Author:     Fasion Chan
 * @contact:   fasionchan@gmail.com
 * @version:   $Id$
 *
 * Description:
 *
 * Changelog:
 *
 **/

package unit

const (
    KB = 1000
    MB = 1000 * KB
    GB = 1000 * MB

    KiB = 1024
    MiB = 1024 * KiB
    GiB = 1024 * MiB
)

type Bytes uint64

func (self Bytes) ToBit() uint64 {
    return uint64(self) * 8
}

func (self Bytes) ToKB() uint64 {
    return uint64(self) / KB
}

func (self Bytes) ToKBf() float64 {
    return float64(self) / KB
}

func (self Bytes) ToMB() uint64 {
    return uint64(self) / MB
}

func (self Bytes) ToMBf() float64 {
    return float64(self) / MB
}

func (self Bytes) ToGB() uint64 {
    return uint64(self) / GB
}

func (self Bytes) ToGBf() float64 {
    return float64(self) / GB
}

func (self Bytes) ToKiB() uint64 {
    return uint64(self) / KiB
}

func (self Bytes) ToKiBf() float64 {
    return float64(self) / KiB
}

func (self Bytes) ToMiB() uint64 {
    return uint64(self) / MiB
}

func (self Bytes) ToMiBf() float64 {
    return float64(self) / MiB
}

func (self Bytes) ToGiB() uint64 {
    return uint64(self) / GiB
}

func (self Bytes) ToGiBf() float64 {
    return float64(self) / GiB
}
