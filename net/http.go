/*
 * Author: fasion
 * Created time: 2019-06-25 18:36:42
 * Last Modified by: fasion
 * Last Modified time: 2019-06-25 18:39:15
 */

package net

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, dst string) (error) {
	out, err := os.OpenFile(dst, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0700)
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
