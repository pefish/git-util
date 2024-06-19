package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
	go_prompt "github.com/pefish/go-prompt"
)

type MergeMainCommand struct {
}

func NewMergeMainCommand() *MergeMainCommand {
	return &MergeMainCommand{}
}

func (dc *MergeMainCommand) Config() interface{} {
	return nil
}

func (dc *MergeMainCommand) Data() interface{} {
	return nil
}

func (dc *MergeMainCommand) Init(command *commander.Commander) error {
	return nil
}

func (dc *MergeMainCommand) OnExited(command *commander.Commander) error {
	return nil
}

func (dc *MergeMainCommand) Start(command *commander.Commander) error {
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
git commit -m "%s"
git pull
git push

current_branch=$(git rev-parse --abbrev-ref HEAD)

git checkout main

git merge ${current_branch}
git pull
git push origin main
git checkout ${current_branch}
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
