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
		Action: func(cCtx *cli.Context) error {
			// args
			path := "./"
			if cCtx.NArg() == 1 {
				path = cCtx.Args().Get(0)
			} else if cCtx.NArg() >= 2 {
				return fmt.Errorf(fmt.Sprintf("Expected 0-1 arguments, but got %v", cCtx.NArg()))
			}
			// scan
			scanResult, err := internal.Scan(path)
			if err != nil {
				return err
			}
			// print scan result in YAML
			yaml, err := internal.Yaml(scanResult)
			if err != nil {
				return err
			}
			fmt.Println(yaml)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
