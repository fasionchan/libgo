/**
 * FileName:   parse.go
 * Author:     Fasion Chan
 * @contact:   fasionchan@gmail.com
 * @version:   $Id$
 *
 * Description:
 *
 * Changelog:
 *
 **/

package parse

import (
    "strconv"
)

func ParseUint64Slice(items []string, values []uint64) ([]uint64, error) {
    if values == nil {
        values = make([]uint64, 0, len(items))
    }

    for _, item := range items {
        value, err := strconv.ParseUint(item, 10, 64)
        if err != nil {
            return nil, err
        }

        values = append(values, value)
    }

    return values, nil
}
