package main

import "my-cni/ipam"

func main() {
	println(ipam.RetrieveNextIpFromAIp("10.123.1.0", 16, 24))
}
