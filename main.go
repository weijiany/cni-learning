package main

import (
	"encoding/json"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
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
