package main

import (
	"github.com/urfave/cli"
	"golang-boilerplate/cmd"
	"golang-boilerplate/setting"
	"log"
	"os"
)

func init() {
	_ = setting.InitMysql()
}

func main() {

	app := cli.NewApp()

	app.Name = "Golang Boilerplate"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		cmd.Start,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}