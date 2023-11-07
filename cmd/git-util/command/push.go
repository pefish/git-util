package command

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/git-util/pkg/global"
	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
)

type PushCommand struct {
}

func NewPushCommand() *PushCommand {
	return &PushCommand{}
}

func (dc *PushCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	return nil
}

func (dc *PushCommand) Init(data *commander.StartData) error {
	err := go_config.ConfigManagerInstance.Unmarshal(&global.GlobalConfig)
	if err != nil {
		return err
	}

	return nil
}

func (dc *PushCommand) OnExited(data *commander.StartData) error {
	return nil
}

func (dc *PushCommand) Start(data *commander.StartData) error {
	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
git add .
git commit -m "%s"
git push
`,
		os.Args[2],
	)
	cmd := exec.Command("bash", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
