package cmd

import (
	"ubctl/src/cmd/clean"
	"ubctl/src/cmd/shell"
	"ubctl/src/cmd/version"
	"ubctl/src/cmd/vm"

	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	version.Commands,
	shell.Commands,
	clean.Commands,
	vm.Commands,
}
