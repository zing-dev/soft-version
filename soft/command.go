package soft

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"time"
)

const (
	filename = "version.json"
)

type Cli struct {
	*cli.App
	Src  []byte
	Soft *Soft
}

//配置文件初始化
func initJson() error {
	_, err := os.Open("version.json")
	if os.IsNotExist(err) {
		file, err := os.Create("version.json")
		if err != nil {
			return errors.New("创建版本配置文件失败")
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
					CreatedAt: fmt.Sprintf("%s", time.Now().Format("2006.01.02 15:04:05")),
				},
			},
			Copyright: "All rights reserved",
		}
		content, err := json.MarshalIndent(soft, "", "  ")
		if err != nil {
			return errors.New("")
		}
		_, err = file.Write(content)
		if err != nil {
			return errors.New("写入版本配置文件失败")
		}
		fmt.Println("初始化版本控制文件成功")
		return err
	}
	return errors.New("版本配置文件已经存在")
}

func (c *Cli) Run(arguments []string) error {
	if len(c.Src) == 0 {
		err := c.init()
		if err != nil {
			return err
		}
	}
	c.Soft = new(Soft)
	err := json.Unmarshal(c.Src, &c.Soft)
	if err != nil {
		return err
	}
	c.App.Commands = append(c.App.Commands, &cli.Command{
		Name:        "init",
		Aliases:     nil,
		Usage:       "init json version",
		UsageText:   "cmd init",
		Description: "初始化版本配置JSON文件",
		Action: func(context *cli.Context) error {
			return c.init()
		},
		HelpName: "init",
	}, &cli.Command{
		Name:        "build",
		Aliases:     nil,
		Usage:       "build json version",
		UsageText:   "cmd build",
		Description: "Build版本配置JSON文件",
		Action: func(context *cli.Context) error {
			return c.build()
		},
		HelpName: "build",
	}, &cli.Command{
		Name:        "info",
		Usage:       "this is soft info",
		UsageText:   "cmd info",
		Description: "软件信息",
		Action: func(context *cli.Context) error {
			fmt.Println(c.Soft.Info())
			return nil
		},
		HelpName: "info",
	}, &cli.Command{
		Name:        "version",
		Usage:       "this is soft version",
		UsageText:   "cmd version",
		Description: "软件信息",
		Action: func(context *cli.Context) error {
			fmt.Println(c.Soft.SimpleVersion())
			return nil
		},
		HelpName: "version",
	}, &cli.Command{
		Name:        "full-version",
		Usage:       "this is soft full-version",
		UsageText:   "cmd full-version",
		Description: "软件信息",
		Action: func(context *cli.Context) error {
			fmt.Println(c.Soft.FullVersion())
			return nil
		},
		HelpName: "full-version",
	}, &cli.Command{
		Name:        "json",
		Usage:       "this is soft info",
		UsageText:   "cmd json",
		Description: "软件信息json",
		Action: func(context *cli.Context) error {
			fmt.Println(string(c.Src))
			return nil
		},
		HelpName: "json",
	}, &cli.Command{
		Name:        "list",
		Usage:       "this is soft list info",
		UsageText:   "cmd list",
		Description: "软件日志列表",
		Action: func(context *cli.Context) error {
			fmt.Println(c.Soft.List())
			return nil
		},
		HelpName: "list",
	}, &cli.Command{
		Name:        "json",
		Usage:       "this is soft info",
		UsageText:   "cmd json",
		Description: "软件信息json",
		Action: func(context *cli.Context) error {
			fmt.Println(string(c.Src))
			return nil
		},
		HelpName: "json",
	})
	return c.App.Run(arguments)
}

func (c *Cli) build() error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c.Soft)
	if err != nil {
		return nil
	}
	hash, _ := Md5FileStr()
	c.Soft.Version[0].Hash = hash
	c.Soft.Version[0].CreatedAt = time.Now().Format("2006.01.02 15:04:05")
	content, err := json.MarshalIndent(c.Soft, "", "  ")
	if err != nil {
		return errors.New(fmt.Sprintf("转换失败: %s", err))
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	_, err = file.Write(content)
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
				CreatedAt: fmt.Sprintf("%s", time.Now().Format("2006.01.02 15:04:05")),
			},
		},
		Copyright: "All rights reserved",
	}
	content, err := json.MarshalIndent(soft, "", "  ")
	if err != nil {
		return errors.New(fmt.Sprintf("转换失败: %s", err))
	}
	_, err = file.Write(content)
	if err != nil {
		return errors.New(fmt.Sprintf("写入版本配置文件: %s", err))
	}
	return nil
}
