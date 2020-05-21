/*
 * Author: fasion
 * Created time: 2020-04-08 09:27:37
 * Last Modified by: fasion
 * Last Modified time: 2020-04-08 09:27:52
 */

package timing

import "time"

func ParseTime(layouts []string, value string) (t time.Time, err error) {
	for _, layout := range layouts {
		if t, err = time.Parse(layout, value); err == nil {
			return
		}
	}

	return
}
