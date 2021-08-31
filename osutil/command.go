/*
 * Author: fasion
 * Created time: 2021-08-31 09:38:04
 * Last Modified by: fasion
 * Last Modified time: 2021-08-31 09:45:21
 */

package osutil

type CommandRoutine = func(string, []string)

type CommandRoutineMapping map[string]CommandRoutine

func (mapping CommandRoutineMapping) Dispatch(args []string) {
	if len(args) == 0 {
		return
	}

	entrypoint := args[0]
	if routine, ok := mapping[entrypoint]; ok && routine != nil {
		routine(entrypoint, args[1:])
	}
}
