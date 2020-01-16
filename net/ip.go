/*
 * Author: fasion
 * Created time: 2019-05-31 15:13:09
 * Last Modified by: fasion
 * Last Modified time: 2019-06-24 17:44:20
 */

package net

import (
	"net"
)

func GetIpv4Addr() (string, error) {
	srcAddr := net.UDPAddr{
		IP: net.IPv4zero,
		Port: 0,
	}
	dstAddr := net.UDPAddr{
		IP: net.ParseIP("8.8.8.8"),
		Port: 8,
	}

	conn, err := net.DialUDP("udp", &srcAddr, &dstAddr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr()
	host, _, err := net.SplitHostPort(localAddr.String())
	if err != nil {
		return "", err
	}


	return host, nil
}
