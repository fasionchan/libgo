/*
 * Author: fasion
 * Created time: 2021-10-22 13:54:11
 * Last Modified by: fasion
 * Last Modified time: 2022-03-22 12:42:24
 */

package goutil

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

var Year2050 = time.Date(2050, 1, 1, 0, 0, 0, 0, time.Now().Local().Location())
var Year3000 = time.Date(3000, 1, 1, 0, 0, 0, 0, time.Now().Local().Location())
var Year5000 = time.Date(5000, 1, 1, 0, 0, 0, 0, time.Now().Local().Location())

func DayOf(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func Today() time.Time {
	return DayOf(time.Now().Local())
}

func ReadableDuration(d time.Duration) string {
	return ""
}

type IntraDayTime time.Duration

func ParseIntraDayTime(s string) (IntraDayTime, error) {
	var hours, minutes, seconds, nanoseconds time.Duration

	var parts = strings.Split(s, ".")
	if _, err := fmt.Sscanf(parts[0], "%d:%d:%d", &hours, &minutes, &seconds); err != nil {
		return 0, err
	}

	// parse nanosecond if any
	if len(parts) > 1 {
		if ns := parts[1]; ns != "" {
			if _, err := fmt.Sscanf(ns, "%d", &nanoseconds); err != nil {
				return 0, err
			}
		}
	}

	return IntraDayTime(hours*time.Hour + minutes*time.Minute + seconds*time.Second + nanoseconds*time.Nanosecond), nil
}

func MustParseIntraDayTime(s string) IntraDayTime {
	t, err := ParseIntraDayTime(s)
	if err != nil {
		panic(err)
	}

	return t
}

func ParseFormattedIntraDayTime(layout, value string) (IntraDayTime, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return 0, err
	}

	r := IntraDayTime(time.Hour)*IntraDayTime(t.Hour()) +
		IntraDayTime(time.Minute)*IntraDayTime(t.Minute()) +
		IntraDayTime(time.Second)*IntraDayTime(t.Second()) +
		IntraDayTime(time.Nanosecond)*IntraDayTime(t.Nanosecond())

	return r, nil
}

func MustParseFormattedIntraDayTime(layout, value string) IntraDayTime {
	t, err := ParseFormattedIntraDayTime(layout, value)
	if err != nil {
		panic(err)
	}

	return t
}

func (t IntraDayTime) Duration() time.Duration {
	return time.Duration(t)
}

func (t IntraDayTime) String() string {
	hours, minutes, seconds, nanoseconds := t.Parts()
	return fmt.Sprintf("%02d:%02d:%02d.%09d", hours, minutes, seconds, nanoseconds)
}

func (t IntraDayTime) Format(layout string) string {
	hours, minutes, seconds, nanoseconds := t.Parts()
	return time.Date(0, 0, 0, hours, minutes, seconds, nanoseconds, time.Local).Format(layout)
}

func (t IntraDayTime) Parts() (int, int, int, int) {
	d := time.Duration(t)

	hours := d / time.Hour
	d -= hours * time.Hour

	minutes := d / time.Minute
	d -= minutes * time.Minute

	seconds := d / time.Second
	d -= seconds * time.Second

	return int(hours), int(minutes), int(seconds), int(d)
}

func (t IntraDayTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *IntraDayTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := ParseIntraDayTime(s)
	if err != nil {
		return err
	}

	*t = parsed

	return nil
}
