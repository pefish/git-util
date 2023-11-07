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

type TagCommand struct {
}

func NewTagCommand() *TagCommand {
	return &TagCommand{}
}

func (dc *TagCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	return nil
}

func (dc *TagCommand) Init(data *commander.StartData) error {
	err := go_config.ConfigManagerInstance.Unmarshal(&global.GlobalConfig)
	if err != nil {
		return err
	}

	return nil
}

func (dc *TagCommand) OnExited(data *commander.StartData) error {
	return nil
}

func (dc *TagCommand) Start(data *commander.StartData) error {
	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail
git add .
git commit -m "%s"
git push
git tag -a "%s" -m "%s"
git push origin "%s"
`,
		os.Args[3],
		os.Args[2],
		os.Args[3],
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
