package main

import (
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/util/log"
	"github.com/tian-yuan/iot-common/basic/config"
	z "github.com/tian-yuan/iot-common/plugins/zap"

	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/tian-yuan/CMQ/topic-acl/commands"
	"github.com/tian-yuan/iot-common/basic"
	"github.com/tian-yuan/iot-common/util"
)

func initLogger() {
	source := env.NewSource()
	basic.Init(
		config.WithSource(source),
	)
	z.Init("iot", "topicacl", "config", "log")
	log.SetLogger(z.GetLogger())
}

func writePid() {
	pathName, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("get executable path failed : %s\n", err))
		return
	}
	index := strings.LastIndex(pathName, "/")
	path := string(pathName[0:index])
	executableName := string(pathName[index+1 : len(pathName)])
	pidFile := path + "/pid/" + executableName + ".pid"
	if err = util.WritePidFile(pidFile); err != nil {
		panic(fmt.Errorf("write pid file failed, pid file : %s, err : %s\n", pidFile, err))
	}
	log.Infof("write pid file %s success.", pidFile)
}

func main() {
	initLogger()
	runtime.GOMAXPROCS(runtime.NumCPU())

	writePid()
	commands.Execute()

	stopCh := util.SetupSignalHandler()
	<-stopCh

	log.Infof("topic acl stop.")
	commands.Stop()
}
