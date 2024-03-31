package ipam

import (
	"math/big"
	"net"
)

func nextIP(ip net.IP, bits int) net.IP {
	ipb := big.NewInt(0).SetBytes(ip)
	ipb.Add(ipb, big.NewInt(1<<(32-bits)))

	b := ipb.Bytes()
	b = append(make([]byte, len(ip)-len(b)), b...)
	return b
}

func RetrieveNextIpFromAIp(aip string, ones, cidrBits int) string {
	ip := net.ParseIP(aip).To4()
	ip = ip.Mask(net.CIDRMask(ones, 32))
	return nextIP(ip, cidrBits).String()
}
