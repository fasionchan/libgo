/*
 * Author: fasion
 * Created time: 2022-02-09 08:41:38
 * Last Modified by: fasion
 * Last Modified time: 2022-02-09 09:18:39
 */

package netutil

import (
	"net/http"
	"testing"

	"github.com/fasionchan/libgo/goutil"
)

func headerEqual(h1, h2 http.Header) bool {
	for name, value1 := range h1 {
		value2 := h2[name]
		if !goutil.StringSlice(value1).Sort().Equal(goutil.StringSlice(value2).Sort()) {
			return false
		}
	}

	for name, value2 := range h2 {
		value1 := h1[name]
		if !goutil.StringSlice(value1).Sort().Equal(goutil.StringSlice(value2).Sort()) {
			return false
		}
	}

	return true
}

func TestHttpHeaderTextParse(t *testing.T) {
	cases := []struct {
		text   string
		header http.Header
	}{
		{
			text:   "",
			header: nil,
		},
		{
			text: "Content-Type: application/json\nAccept: */*",
			header: http.Header{
				"Content-Type": []string{
					"application/json",
				},
				"Accept": []string{
					"*/*",
				},
			},
		},
		{
			text: "Content-Type: application/json\nAccept: */*\n",
			header: http.Header{
				"Content-Type": []string{
					"application/json",
				},
				"Accept": []string{
					"*/*",
				},
			},
		},
		{
			text: "\n\n\nContent-Type: application/json\n\n\nAccept: */*\n\n\n",
			header: http.Header{
				"Content-Type": []string{
					"application/json",
				},
				"Accept": []string{
					"*/*",
				},
			},
		},
		{
			text: "Cookie: C=D\nContent-Type: application/json\nAccept: */*\nCookie: A=B",
			header: http.Header{
				"Content-Type": []string{
					"application/json",
				},
				"Accept": []string{
					"*/*",
				},
				"Cookie": []string{
					"A=B",
					"C=D",
				},
			},
		},
	}

	for _, testCase := range cases {
		header := HttpHeaderText(testCase.text).Parse()
		if !headerEqual(header, testCase.header) {
			t.Errorf("case failed: %s should be %v but got %v", testCase.text, testCase.header, header)
		}
	}
}
