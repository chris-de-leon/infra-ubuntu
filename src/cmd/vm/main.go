package vm

import (
	"ubctl/src/cmd/vm/apply"
	"ubctl/src/cmd/vm/rollback"

	"github.com/urfave/cli/v2"
)

var Commands = &cli.Command{
	Name:  "vm",
	Usage: "Commands for managing Ubuntu VMs",
	Subcommands: []*cli.Command{
		rollback.Commands,
		apply.Commands,
		initCmd,
		undoCmd,
		runCmd,
	},
}
