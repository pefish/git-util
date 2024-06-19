package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
	go_prompt "github.com/pefish/go-prompt"
)

type TagCommand struct {
}

func NewTagCommand() *TagCommand {
	return &TagCommand{}
}

type TagCommandConfigType struct {
}

type TagCommandDataType struct {
}

var TagCommandConfig TagCommandConfigType
var TagCommandData TagCommandDataType

func (dc *TagCommand) Init(command *commander.Commander) error {
	return nil
}

func (dc *TagCommand) Config() interface{} {
	return &TagCommandConfig
}

func (dc *TagCommand) Data() interface{} {
	return &TagCommandData
}

func (dc *TagCommand) OnExited(command *commander.Commander) error {
	return nil
}

func (dc *TagCommand) Start(command *commander.Commander) error {
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

	tag, isExit := go_prompt.PromptInstance.Input(
		"Please input tag.",
		nil,
	)
	if isExit {
		return nil
	}
	if tag == "" {
		go_logger.Logger.InfoFRaw("Error: required 'tag' not specified.")
		return nil
	}

	script := fmt.Sprintf(
		`
#!/bin/bash
set -euxo pipefail

git config pull.rebase false
git add .
git commit -m "%s"
git pull
git push
git tag -a "%s" -m "%s"
git push origin "%s"
`,
		comment,
		tag,
		comment,
		tag,
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
