/*
 * Author: fasion
 * Created time: 2021-10-25 09:46:57
 * Last Modified by: fasion
 * Last Modified time: 2021-10-25 10:17:01
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
