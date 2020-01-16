/**
 * FileName:   slice.go
 * Author:     Fasion Chan
 * @contact:   fasionchan@gmail.com
 * @version:   $Id$
 *
 * Description:
 *
 * Changelog:
 *
 **/

package math

import (
    "fmt"
    "github.com/fasionchan/libgo/copy"
)

var _ = fmt.Println

func Uint64SliceSub(a []uint64, b []uint64, inplace bool) ([]uint64) {
    var fr []uint64
    if inplace {
        fr = a
    } else {
        fr = copy.DeepCopyUint64Slice(a)
    }

    limit := len(fr)

    for i, value := range b {
        if i >= limit {
            break
        }

        fr[i] -= value
    }

    return fr
}

func Uint64SliceSub2D(a []([]uint64), b []([]uint64), inplace bool) ([]([]uint64)) {
    var fr []([]uint64)
    if inplace {
        fr = a
    } else {
        fr = copy.DeepCopyUint64Slice2D(a)
    }

    limit := len(fr)

    for i, values := range b {
        if i >= limit {
            break
        }

        Uint64SliceSub(fr[i], values, true)
    }

    return fr
}

func SumUint64Slice(values []uint64) (uint64) {
    var sum uint64 = 0
    for _, value := range values {
        sum += value
    }
    return sum
}
