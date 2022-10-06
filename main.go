package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mys1024/imgi/internal"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "imgi",
		Usage: "prints images' information",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Value:   "yaml",
				Usage:   "output format",
				Aliases: []string{"f"},
			},
		},
		Action: func(cCtx *cli.Context) error {
			// flags
			format := cCtx.String("format")
			formatter := internal.Yaml
			switch format {
			case "yaml":
				formatter = internal.Yaml
			case "toml":
				formatter = internal.Toml
			case "json":
				formatter = internal.Json
			default:
				return fmt.Errorf(fmt.Sprintf("Wrong flag (format) value: %v", format))
			}
			// args
			dir := "./"
			if cCtx.NArg() == 1 {
				dir = cCtx.Args().Get(0)
			} else if cCtx.NArg() >= 2 {
				return fmt.Errorf(fmt.Sprintf("Expected 0-1 arguments, but got %v", cCtx.NArg()))
			}
			// scan
			scanResult, err := internal.Scan(dir)
			if err != nil {
				return err
			}
			// output
			output, err := formatter(scanResult)
			if err != nil {
				return err
			}
			fmt.Println(output)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
