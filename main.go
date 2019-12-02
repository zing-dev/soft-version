package main

import (
	soft "github.com/zhangrxiang/soft-version/src"
)

func main() {
	command := soft.NewCommand()
	command.Parse()
}
