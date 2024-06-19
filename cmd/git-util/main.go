package main

import (
	"github.com/pefish/git-util/cmd/git-util/command"
	"github.com/pefish/git-util/version"
	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName+" is a template.")
	commanderInstance.RegisterSubcommand("push", &commander.SubcommandInfo{
		Desc:       "Push code.",
		Args:       nil,
		Subcommand: command.NewPushCommand(),
	})
	commanderInstance.RegisterSubcommand("tag", &commander.SubcommandInfo{
		Desc:       "Tag code.",
		Args:       nil,
		Subcommand: command.NewTagCommand(),
	})
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}
