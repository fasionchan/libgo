/*
 * Author: fasion
 * Created time: 2019-07-12 16:21:07
 * Last Modified by: fasion
 * Last Modified time: 2019-07-12 16:22:22
 */

package sorting

import (
	"strings"
)

func CompareVersionString(v1, v2, s string) (int) {
	return CompareStringSlice(
		strings.Split(v1, s),
		strings.Split(v2, s),
	)
}
