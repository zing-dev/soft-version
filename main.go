package main

import (
	"flag"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	soft "soft-version/src"
)

func main() {
	//v := &version.SoftVersion{}
	//v.Parse()
	//test()
	//test2()
	//soft.NewCommand()
	command := soft.NewCommand()
	command.Parse()

}

func test2() {
	app := kingpin.New("DTS", "无锡亚天科技DTS平台软件")
	app.Version("1.0.1").Author("无锡亚天光电")
	app.Command("run", "运行DTS平台").Action(func(context *kingpin.ParseContext) error {
		fmt.Println("run.....")
		return nil
	})
	app.Command("version", "DTS平台软件版本").Action(func(context *kingpin.ParseContext) error {
		fmt.Println("1.1.1")
		return nil
	})
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
func test() {
	log.Println(len(os.Args))
	log.Println(os.Args[0])
	if len(os.Args) > 1 {
		flag.Parse()
		name := flag.String("name", "zing", "name")
		log.Println(*name)
	}
}
