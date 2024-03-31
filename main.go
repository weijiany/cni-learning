package main

import (
	"encoding/json"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
	"my-cni/ipam"
	"my-cni/utils"
)

type PluginConfig struct {
	types.NetConf

	Bridge string `json:"bridge"`
	Subnet string `json:"subnet"`
}

func cmdAdd(args *skel.CmdArgs) error {
	pluginConfig := &PluginConfig{}
	err := json.Unmarshal(args.StdinData, pluginConfig)
	if err != nil {
		panic(err)
	}
	am := ipam.GetIpAM(pluginConfig.Subnet)
	am.RecordHostIpIfNotExist()
	return nil
}

func cmdCheck(args *skel.CmdArgs) error {
	utils.WriteLog(args.ContainerID)
	utils.WriteLog(args.Netns)
	utils.WriteLog(args.IfName)
	utils.WriteLog(args.Args)
	utils.WriteLog(args.Path)
	return nil
}

func cmdDel(args *skel.CmdArgs) error {
	utils.WriteLog(args.ContainerID)
	utils.WriteLog(args.Netns)
	utils.WriteLog(args.IfName)
	utils.WriteLog(args.Args)
	utils.WriteLog(args.Path)
	return nil
}

func main() {
	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.All, bv.BuildString("my-cni"))
}

/**
IPAM need to support the following features:
1. record the IP that has been used by node
2. get the unused IP based on one node
3. retrieve the node IP
4. get all node IPs
*/
