package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
	"my-cni/utils"
)

func cmdAdd(args *skel.CmdArgs) error {
	utils.WriteLog(args.ContainerID)
	utils.WriteLog(args.Netns)
	utils.WriteLog(args.IfName)
	utils.WriteLog(args.Args)
	utils.WriteLog(args.Path)
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
