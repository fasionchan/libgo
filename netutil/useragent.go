/*
 * Author: fasion
 * Created time: 2021-08-17 09:01:29
 * Last Modified by: fasion
 * Last Modified time: 2021-08-17 09:08:12
 */

package netutil

import "strings"

type UserAgent string

func (ua UserAgent) Raw() string {
	return string(ua)
}

func (ua UserAgent) ContainAnyKeywords(keywords ...string) bool {
	raw := strings.ToLower(ua.Raw())
	for _, keyword := range keywords {
		if strings.Contains(raw, keyword) {
			return true
		}
	}

	return false
}

func (ua UserAgent) IsPc() bool {
	return ua.ContainAnyKeywords(
		"windows",
		"macintosh",
		"x86",
		"x86_64",
		"x64",
	)
}

func (ua UserAgent) IsWxWork() bool {
	return ua.ContainAnyKeywords(
		"wxwork",
	)
}

func (ua UserAgent) IsWeixin() bool {
	return ua.ContainAnyKeywords(
		"weixin",
	)
}

func (ua UserAgent) IsWxWorkOrWeixin() bool {
	return ua.IsWxWork() || ua.IsWeixin()
}

func (ua UserAgent) IsWindows() bool {
	return ua.ContainAnyKeywords(
		"windows",
	)
}

func (ua UserAgent) IsMac() bool {
	return ua.ContainAnyKeywords(
		"macintosh",
	)
}
