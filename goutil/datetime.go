/*
 * Author: fasion
 * Created time: 2021-10-22 13:54:11
 * Last Modified by: fasion
 * Last Modified time: 2021-10-25 10:15:20
 */

package goutil

import (
	"encoding/json"
	"fmt"
	"time"
)

func DayOf(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func Today() time.Time {
	return DayOf(time.Now().Local())
}

type IntraDayTime time.Duration

func (t IntraDayTime) Duration() time.Duration {
	return time.Duration(t)
}

func (t IntraDayTime) MarshalJSON() ([]byte, error) {
	d := time.Duration(t)

	hours := d / time.Hour
	d -= hours * time.Hour

	minutes := d / time.Minute
	d -= minutes * time.Minute

	seconds := d / time.Second
	d -= seconds * time.Second

	return json.Marshal(fmt.Sprintf("%02d:%02d:%02d.%09d", hours, minutes, seconds, d))
}

func (t *IntraDayTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	var hours, minutes, seconds, nanoseconds time.Duration
	if _, err := fmt.Sscanf(s, "%d:%d:%d.%d", &hours, &minutes, &seconds, &nanoseconds); err != nil {
		return err
	}

	*t = IntraDayTime(hours*time.Hour + minutes*time.Minute + seconds*time.Second + nanoseconds*time.Nanosecond)

	return nil
}
