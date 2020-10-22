package cli_demo

import (
	"os"

	"github.com/urfave/cli/v2"
)

// SingleCMD -
func SingleCMD() {
	(&cli.App{}).Run(os.Args)
}
