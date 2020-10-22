package cli_demo

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Boom -
func Boom() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(context *cli.Context) error {
			fmt.Println("Boooooooom!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
