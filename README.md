# 软件版本命令控制显示

## Example
first use the library such as:
```go
package main

import (
	soft "github.com/zhangrxiang/soft-version/src"
)

func main() {
	soft.NewCommand().Parse()
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
  "name": "name",
  "alias": "alias",
  "author": "author",
  "version": {
    "version": "0.0.0",
    "log": "init",
    "status": "Base"
  },
  "copyright": "Copyright",
}
```
after go on
```shell script
$ ./[You Soft] build
$ go build
$ ./[You Soft] help
author

Flags:
  --help  Show context-sensitive help (also try --he
lp-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  run
    运行soft-version

  version
    soft-version软件版本

  full-version
    soft-version软件版本全称

  info
    soft-version软件版本信息

  build
    开发编译[软件开发者专用]

  init
    初始化软件版本配置文件[软件开发者专用]

```