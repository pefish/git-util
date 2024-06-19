package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pefish/go-commander"
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
