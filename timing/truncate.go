/*
 * Author: fasion
 * Created time: 2019-06-18 09:39:46
 * Last Modified by: fasion
 * Last Modified time: 2019-06-18 09:40:07
 */

package timing

import (
	"time"
)

func TruncateLocal(t time.Time, d time.Duration) (time.Time) {
	if d <= time.Minute {
		return t.Truncate(d)
	}

	_, offset := t.Zone()
	step := d.Nanoseconds() / 1000000000

	unix := t.Unix() + int64(offset)
	unix = unix / step * step - int64(offset)

	return time.Unix(unix, 0).In(t.Location())
}
