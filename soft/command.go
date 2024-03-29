package soft

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"sync"
	"time"
)

const (
	filename = "version.json"
)

var (
	once = sync.Once{}
	c    *Cli
)

type Cli struct {
	*cli.App
	Src  []byte
	Soft *Soft
}

// NewCli 实例化 Cli
func NewCli(app *cli.App, src []byte) *Cli {
	once.Do(func() {
		c = &Cli{
			App: app,
			Src: src,
		}
	})
	return c
}

// GetCli 获取 Cli
func GetCli() *Cli {
	return c
}

func (c *Cli) Run(arguments []string) error {
	if len(c.Src) == 0 {
		err := c.init()
		if err != nil {
			return errors.New(fmt.Sprintf("init err: %s", err))
		}
	}
	c.Soft = new(Soft)
	err := json.Unmarshal(c.Src, &c.Soft)
	if err != nil {
		return errors.New(fmt.Sprintf("unmarshal err: %s", err))
	}
	c.App.Commands = append(c.App.Commands, &cli.Command{
		Name:        "dev",
		Usage:       "仅供开发者使用的命令",
		UsageText:   "dev",
		Description: "仅供开发者使用的命令",
		Subcommands: []*cli.Command{
			{
				Name:        "init",
				Usage:       "初始化版本文件",
				UsageText:   "cmd init",
				Description: "初始化版本配置JSON文件",
				Action: func(context *cli.Context) error {
					return c.init()
				},
				HelpName: "init",
			}, {
				Name:        "build",
				Usage:       "更新版本文件",
				UsageText:   "cmd build",
				Description: "Build版本配置JSON文件",
				Action: func(context *cli.Context) error {
					return c.build()
				},
				HelpName: "build",
			},
		},
	}, &cli.Command{
		Name:        "about",
		Usage:       "关于软件的详情",
		UsageText:   "version",
		Description: "使用 about 子命令 查看软件信息",
		Subcommands: []*cli.Command{
			{
				Name:        "info",
				Usage:       "this is soft info",
				UsageText:   "cmd info",
				Description: "软件信息",
				Action: func(context *cli.Context) error {
					fmt.Println(c.Soft.Info())
					return nil
				},
				HelpName: "info",
			}, {
				Name:        "version",
				Usage:       "this is soft version",
				UsageText:   "cmd version",
				Description: "软件信息",
				Action: func(context *cli.Context) error {
					fmt.Println(c.Soft.SimpleVersion())
					return nil
				},
				HelpName: "version",
			}, {
				Name:        "full-version",
				Usage:       "this is soft full-version",
				UsageText:   "cmd full-version",
				Description: "软件信息",
				Action: func(context *cli.Context) error {
					fmt.Println(c.Soft.FullVersion())
					return nil
				},
				HelpName: "full-version",
			}, {
				Name:        "json",
				Usage:       "this is soft info",
				UsageText:   "cmd json",
				Description: "软件信息json",
				Action: func(context *cli.Context) error {
					fmt.Println(string(c.Src))
					return nil
				},
				HelpName: "json",
			}, {
				Name:        "list",
				Usage:       "this is soft list info",
				UsageText:   "cmd list",
				Description: "软件日志列表",
				Action: func(context *cli.Context) error {
					fmt.Println(c.Soft.List())
					return nil
				},
				HelpName: "list",
			}, {
				Name:        "json",
				Usage:       "this is soft info",
				UsageText:   "cmd json",
				Description: "软件信息json",
				Action: func(context *cli.Context) error {
					fmt.Println(string(c.Src))
					return nil
				},
				HelpName: "json",
			},
		}},
	)
	return c.App.Run(arguments)
}

func (c *Cli) build() error {
	data, err := os.ReadFile(filename)
	err = json.Unmarshal(data, &c.Soft)
	if err != nil {
		return nil
	}
	hash, _ := Md5FileStr()
	c.Soft.Version[0].Hash = hash
	c.Soft.Version[0].CreatedAt = time.Now().Format(time.DateTime)
	data, err = json.MarshalIndent(c.Soft, "", "  ")
	if err != nil {
		return errors.New(fmt.Sprintf("转换失败: %s", err))
	}
	file, err := os.Create(filename)
	err = json.Unmarshal(data, &c.Soft)
	if err != nil {
		return nil
	}
	_, err = file.Write(data)
	if err != nil {
		return errors.New(fmt.Sprintf("写入版本配置文件: %s", err))
	}
	return nil
}

func (c *Cli) init() error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c.Soft)
	if err == nil {
		return nil
	}
	hash, _ := Md5FileStr()
	soft := &Soft{
		Name:   "xx-软件",
		Alias:  "别名",
		Author: "作者",
		Version: []Version{
			{
				Tag:       "0.0.1",
				Log:       "init",
				Status:    Base,
				Hash:      hash,
				CreatedAt: fmt.Sprintf("%s", time.Now().Format(time.DateTime)),
			},
		},
		Copyright: "All rights reserved",
	}
	content, err := json.MarshalIndent(soft, "", "  ")
	if err != nil {
		return errors.New(fmt.Sprintf("转换失败: %s", err))
	}
	_, _ = file.Seek(0, io.SeekStart)
	_, err = file.Write(content)
	if err != nil {
		return errors.New(fmt.Sprintf("写入版本配置文件: %s", err))
	}
	return nil
}
