package ipam

import (
	"my-cni/etcdutils"
	"os"
	"strconv"
	"strings"
)

type IP struct {
	address string
	prefix  int
}

type IpAM struct {
	hostName   string
	etcdClient *etcdutils.Client
	ip         IP
}

func GetIpAM(subnet string) *IpAM {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	address := strings.Split(subnet, "/")
	prefix, err := strconv.Atoi(address[1])
	if err != nil {
		panic(err)
	}

	return &IpAM{
		hostName:   hostname,
		etcdClient: etcdutils.GetClient(),
		ip:         IP{address: address[0], prefix: prefix},
	}
}

func (ipam *IpAM) RecordHostIpIfNotExist() {
	hostIp := ipam.recordSelf()
	ipam.recordAllNodes(hostIp)
}

func (ipam *IpAM) recordSelf() string {
	var ipAddress string
	if tmpIp := ipam.etcdClient.Get(ipam.hostName, "ip"); tmpIp != "" {
		ipAddress = tmpIp
	} else {
		allNodeIps := strings.Split(ipam.etcdClient.Get("all-node-ips"), ",")
		ipAddress = allNodeIps[len(allNodeIps)-1]
	}

	hostIp := RetrieveNextIpFromAIp(ipAddress, ipam.ip.prefix, ipam.ip.prefix+8)
	ipam.etcdClient.Put(hostIp, ipam.hostName, "ip")
	return hostIp
}

func (ipam *IpAM) recordAllNodes(hostIp string) {
	allNodeIps := ""
	if tmpIp := ipam.etcdClient.Get("all-node-ips"); tmpIp != "" {
		allNodeIps = tmpIp
	}
	if !strings.Contains(allNodeIps, hostIp) {
		ips := []string{allNodeIps, hostIp}
		if allNodeIps == "" {
			ips = ips[1:]
		}
		ipam.etcdClient.Put(strings.Join(ips, ","), "all-node-ips")
	}
}
