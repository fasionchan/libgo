/*
 * Author: fasion
 * Created time: 2019-07-12 16:10:53
 * Last Modified by: fasion
 * Last Modified time: 2019-07-12 16:20:23
 */

package sorting

func CompareStringSlice(s1, s2 []string) (int) {
	n1, n2 := len(s1), len(s2)
	var i, j int

	for (i < n1) {
		if j == n2 {
			return 1;
		}

		if s1[i] > s2[j] {
			return 1;
		}

		if s1[i] < s2[j] {
			return -1;
		}

		i += 1
		j += 1
	}

	if j < n2 {
		return -1;
	}

	return 0;
}
