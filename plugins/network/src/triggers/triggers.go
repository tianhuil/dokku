package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dokku/dokku/plugins/common"
	"github.com/dokku/dokku/plugins/network"
)

// main entrypoint to all triggers
func main() {
	parts := strings.Split(os.Args[0], "/")
	trigger := parts[len(parts)-1]
	flag.Parse()

	var err error
	switch trigger {
	case "docker-args-process-build":
		appName := flag.Arg(0)
		err = network.TriggerDockerArgsProcess(appName)
	case "docker-args-process-deploy":
		appName := flag.Arg(0)
		err = network.TriggerDockerArgsProcess(appName)
	case "docker-args-process-run":
		appName := flag.Arg(0)
		err = network.TriggerDockerArgsProcess(appName)
	case "install":
		err = network.TriggerInstall()
	case "network-build-config":
		appName := flag.Arg(0)
		err = network.BuildConfig(appName)
	case "network-compute-ports":
		appName := flag.Arg(0)
		processType := flag.Arg(1)
		isHerokuishContainer := common.ToBool(flag.Arg(2))
		err = network.TriggerNetworkComputePorts(appName, processType, isHerokuishContainer)
	case "network-config-exists":
		appName := flag.Arg(0)
		err = network.TriggerNetworkConfigExists(appName)
	case "network-get-ipaddr":
		appName := flag.Arg(0)
		processType := flag.Arg(1)
		containerID := flag.Arg(2)
		err = network.TriggerNetworkGetIppaddr(appName, processType, containerID)
	case "network-get-listeners":
		appName := flag.Arg(0)
		processType := flag.Arg(1)
		err = network.TriggerNetworkGetListeners(appName, processType)
	case "network-get-port":
		appName := flag.Arg(0)
		processType := flag.Arg(1)
		containerID := flag.Arg(2)
		isHerokuishContainer := common.ToBool(flag.Arg(3))
		err = network.TriggerNetworkGetPort(appName, processType, containerID, isHerokuishContainer)
	case "network-get-property":
		appName := flag.Arg(0)
		property := flag.Arg(1)
		err = network.TriggerNetworkGetProperty(appName, property)
	case "network-write-ipaddr":
		appName := flag.Arg(0)
		processType := flag.Arg(1)
		containerIndex := flag.Arg(2)
		ip := flag.Arg(3)
		err = network.TriggerNetworkWriteIpaddr(appName, processType, containerIndex, ip)
	case "network-write-port":
		appName := flag.Arg(0)
		processType := flag.Arg(1)
		containerIndex := flag.Arg(2)
		port := flag.Arg(3)
		err = network.TriggerNetworkWritePort(appName, processType, containerIndex, port)
	case "post-app-clone-setup":
		oldAppName := flag.Arg(0)
		newAppName := flag.Arg(1)
		err = network.TriggerPostAppCloneSetup(oldAppName, newAppName)
	case "post-app-rename-setup":
		oldAppName := flag.Arg(0)
		newAppName := flag.Arg(1)
		err = network.TriggerPostAppRenameSetup(oldAppName, newAppName)
	case "post-container-create":
		containerType := flag.Arg(0)
		containerID := flag.Arg(1)
		appName := flag.Arg(2)
		phase := flag.Arg(3)
		processType := flag.Arg(4)
		err = network.TriggerPostContainerCreate(containerType, containerID, appName, phase, processType)
	case "post-create":
		appName := flag.Arg(0)
		err = network.TriggerPostCreate(appName)
	case "post-delete":
		appName := flag.Arg(0)
		err = network.TriggerPostDelete(appName)
	case "core-post-deploy":
		appName := flag.Arg(0)
		err = network.TriggerCorePostDeploy(appName)
	case "report":
		appName := flag.Arg(0)
		err = network.ReportSingleApp(appName, "", "")
	default:
		err = fmt.Errorf("Invalid plugin trigger call: %s", trigger)
	}

	if err != nil {
		common.LogFailWithError(err)
	}
}
