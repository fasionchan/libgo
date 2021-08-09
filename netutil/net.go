/*
 * Author: fasion
 * Created time: 2021-08-09 15:39:06
 * Last Modified by: fasion
 * Last Modified time: 2021-08-09 15:52:23
 */

package netutil

import "net"

var Ipv4Intranets = IpNets{
	{
		IP:   net.ParseIP("10.0.0.0"),
		Mask: net.CIDRMask(8, 32),
	},
	{
		IP:   net.ParseIP("172.16.0.0"),
		Mask: net.CIDRMask(12, 32),
	},
	{
		IP:   net.ParseIP("192.168.0.0"),
		Mask: net.CIDRMask(16, 32),
	},
}

type IpFilter = func(net.IP) bool

type Ips []net.IP

func (ips Ips) Append(more ...net.IP) Ips {
	return append(ips, more...)
}

func (ips Ips) Filter(f IpFilter) Ips {
	result := make(Ips, 0, len(ips))
	for _, ip := range ips {
		if f(ip) {
			result = result.Append(ip)
		}
	}
	return result
}

func (ips Ips) AllMatch(f IpFilter) bool {
	for _, ip := range ips {
		if !f(ip) {
			return false
		}
	}
	return true
}

func (ips Ips) AllInIpv4Intranets() bool {
	return ips.AllMatch(func(ip net.IP) bool {
		return Ipv4Intranets.Contains(ip)
	})
}

type IpNets []*net.IPNet

func (nets IpNets) Contains(ip net.IP) bool {
	for _, net := range nets {
		if net.Contains(ip) {
			return true
		}
	}

	return false
}
