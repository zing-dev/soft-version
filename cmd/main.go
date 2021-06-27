package main

import (
	_ "embed"
	"github.com/urfave/cli/v2"
	"github.com/zing-dev/soft-version/soft"
	"os"
)

//go:embed version.json
var str []byte

func main() {
	app := soft.Cli{App: cli.NewApp(), Src: str}
	app.Run(os.Args)
}
