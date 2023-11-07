package main

import (
	"github.com/pefish/git-util/cmd/git-util/command"
	"github.com/pefish/git-util/version"
	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName+" is a template.")
	commanderInstance.RegisterSubcommand("push", "Push code.", command.NewPushCommand())
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}
