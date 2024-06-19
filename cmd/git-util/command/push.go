package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
	go_prompt "github.com/pefish/go-prompt"
)

type PushCommand struct {
}

type PushCommandConfigType struct {
}

type PushCommandDataType struct {
}

var PushCommandConfig PushCommandConfigType
var PushCommandData PushCommandDataType

func NewPushCommand() *PushCommand {
	return &PushCommand{}
}

func (dc *PushCommand) Config() interface{} {
	return &PushCommandConfig
}

func (dc *PushCommand) Data() interface{} {
	return &PushCommandData
}

func (dc *PushCommand) Init(command *commander.Commander) error {
	return nil
}

func (dc *PushCommand) OnExited(command *commander.Commander) error {
	return nil
}

func (dc *PushCommand) Start(command *commander.Commander) error {
	comment, isExit := go_prompt.PromptInstance.Input(
		"Please input comment.",
		nil,
	)
	if isExit {
		return nil
	}
	if comment == "" {
		go_logger.Logger.InfoFRaw("Error: required 'comment' not specified.")
		return nil
	}

	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail

git config pull.rebase false

git add .
git pull
git commit -m "%s"
git push
`,
		comment,
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
