/**
 * FileName:   dump.go
 * Author:     Fasion Chan
 * @contact:   fasionchan@gmail.com
 * @version:   $Id$
 *
 * Description:
 *
 * Changelog:
 *
 **/

package copy

import (
    "fmt"
)

var _ = fmt.Println

func DeepCopyUint64Slice(src []uint64) ([]uint64) {
    dup := make([]uint64, len(src))
    copy(dup, src)
    return dup
}

func DeepCopyUint64Slice2D(src []([]uint64)) ([]([]uint64)) {
    dup := make([]([]uint64), len(src))
    for i, sub := range src {
        dup[i] = DeepCopyUint64Slice(sub)
    }

    return dup
}
