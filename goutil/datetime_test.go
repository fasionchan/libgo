/*
 * Author: fasion
 * Created time: 2021-10-25 09:46:57
 * Last Modified by: fasion
 * Last Modified time: 2022-03-22 12:39:43
 */

package goutil

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

var IntraDayTimeJsonCases = []struct {
	t IntraDayTime
	j []byte
}{
	{
		t: 0,
		j: []byte(`"00:00:00.000000000"`),
	},
	{
		t: IntraDayTime(1*time.Hour + 2*time.Minute + 3*time.Second + 4*time.Nanosecond),
		j: []byte(`"01:02:03.000000004"`),
	},
}

func TestMarshalIntraDayTime(t *testing.T) {
	var idt IntraDayTime
	err := json.Unmarshal([]byte(`"11:00:00"`), &idt)
	if err != nil {
		t.Error(err)
	}

	for _, data := range IntraDayTimeJsonCases {
		j, err := json.Marshal(data.t)
		if err != nil {
			t.Error(err)
			continue
		}

		if bytes.Compare(j, data.j) != 0 {
			t.Errorf("`%s` and `%s` is not the same", string(j), string(data.j))
		}
	}
}

func TestUnMarshalIntraDayTime(t *testing.T) {
	for _, data := range IntraDayTimeJsonCases {
		var idt IntraDayTime
		if err := json.Unmarshal(data.j, &idt); err != nil {
			t.Error(err)
			continue
		}

		if idt != data.t {
			t.Errorf("%d and %d is not the same", idt, data.t)
		}
	}
}

var IntraDayTimeFormatCases = []struct {
	t IntraDayTime
	f string
	r string
}{
	{
		t: MustParseIntraDayTime("01:02:03.999999999"),
		f: "15:04",
		r: "01:02",
	},
	{
		t: MustParseIntraDayTime("01:02:03.999999999"),
		f: "04:05",
		r: "02:03",
	},
	{
		t: MustParseFormattedIntraDayTime("15-04-05", "01-02-03"),
		f: "15:04:05",
		r: "01:02:03",
	},
}

func TestFormatIntraDayTime(t *testing.T) {
	for _, data := range IntraDayTimeFormatCases {
		r := data.t.Format(data.f)
		if r != data.r {
			t.Errorf("%s and %s is not the same", r, data.r)
		}
	}
}
