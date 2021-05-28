/*
 * Author: fasion
 * Created time: 2021-05-26 10:57:44
 * Last Modified by: fasion
 * Last Modified time: 2021-05-26 11:00:01
 */

package goutil

import (
	"encoding/json"
	"fmt"
)

func PrintJson(args ...interface{}) {
	for i, args := range args {
		fmt.Printf("### %d ###", i)

		result, err := json.MarshalIndent(args, "", "    ")
		if err == nil {
			fmt.Println(string(result))
		} else {
			fmt.Println(err)
		}
	}
}
