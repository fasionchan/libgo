/*
 * Author: fasion
 * Created time: 2019-06-25 18:36:42
 * Last Modified by: fasion
 * Last Modified time: 2022-02-09 08:43:23
 */

package netutil

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type HttpHeaderText string

func (t HttpHeaderText) Parse() http.Header {
	header := make(http.Header)

	for _, line := range strings.Split(string(t), "\n") {
		fields := strings.SplitN(line, ":", 2)
		if len(fields) != 2 {
			continue
		}

		name := strings.TrimSpace(fields[0])
		if name == "" {
			continue
		}

		header[name] = append(header[name], strings.TrimLeft(fields[1], " "))
	}

	return header
}

func DownloadFile(url, dst string) error {
	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return err
	}
	defer out.Close()

	rsps, err := http.Get(url)
	if err != nil {
		return err
	}
	defer rsps.Body.Close()

	if rsps.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", rsps.Status)
	}

	_, err = io.Copy(out, rsps.Body)
	return err
}
