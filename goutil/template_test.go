/*
 * Author: fasion
 * Created time: 2021-11-18 09:34:42
 * Last Modified by: fasion
 * Last Modified time: 2022-03-08 12:05:17
 */

package goutil

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
	"time"
)

var now = time.Now()

var TimeFormatTemplateHelperCases = []struct {
	t      time.Time
	tpl    string
	result string
}{
	{
		t:      time.Time{},
		tpl:    `{{ timefmt . "20060102" "待定" }}`,
		result: "待定",
	},
	{
		t:      now,
		tpl:    fmt.Sprintf(`{{ timefmt . "%s" "待定" }}`, DefaultTimeFormat),
		result: now.Format(DefaultTimeFormat),
	},
	{
		t:      now,
		tpl:    `{{ timefmt . "" "待定" }}`,
		result: now.Format(DefaultTimeFormat),
	},
}

func renderTemplate(tplText string, data interface{}, funcMap template.FuncMap) (string, error) {
	tpl, err := template.New("test").Funcs(funcMap).Parse(tplText)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func TestTimeFormatTempalteHelper(t *testing.T) {
	for _, item := range TimeFormatTemplateHelperCases {
		result, err := renderTemplate(item.tpl, item.t, TemplateHelpers.Native())
		if err != nil {
			t.Error(err)
			continue
		}

		if result != item.result {
			t.Fatalf("`%s` expected but got `%s`", item.result, result)
		}
	}
}

func TestMapHelpers(t *testing.T) {
	result, err := renderTemplate(`{{ $m := makeMap "" (int64 0) }}
{{ setMap $m "test" (addInt64 (index $m "test") 10) }}
{{ printf "%v: %T\n" $m $m }}
{{ jsonify $m }}
{{ $nums := (jsondecode "[0, 1, 1.2, 1.999999999999999999999999999999999]") }}
{{ range $_, $num := $nums }}{{ printf "%v: %T\n" $num $num }}{{ end }}
{{ $i := convert (index $nums 3) "int" }}
{{ printf "%v: %T\n" $i $i }}
{{ addInt 0 1.1 1.999999999999999999999999999999999999 }}
`, nil, TemplateHelpers.Native())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
