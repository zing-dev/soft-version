# 软件版本命令控制显示

## Example
first use the library such as:
```go
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
	app := soft.Cli{App: cli.NewApp(),Src: str}
	app.Run(os.Args)
}


```
then open your terminal:
```shell script
$ go build

$ ./[You Soft] init
```
update version.json file under your project,input your project info
```json
{
  "name": "xx-软件",
  "alias": "别名",
  "author": "作者",
  "version": {
    "version": "0.0.1",
    "log": "init",
    "status": "Base"
  },
  "copyright": "All rights reserved",
  "inherit": true
}
```
after go on
```shell script
$ ./[You Soft] build
$ go build
$ ./[You Soft] help

usage: xx-软件 [<flags>] <command> [<args> ...]

作者

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man)
.

Commands:
  help [<command>...]
    Show help.

  run
    运行xx-软件

  version
    xx-软件版本

  full-version
    xx-软件版本全称

  info
    xx-软件版本信息

  build
    开发编译[软件开发者专用]

  init
    初始化软件版本配置文件[软件开发者专用]

```