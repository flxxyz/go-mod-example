package main

import (
	"github.com/op/go-logging"
	"gopkg.in/urfave/cli.v1"
	"os"
)

var log = logging.MustGetLogger("example")

func main() {
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")
	logging.SetBackend(backend1Leveled, backend2Formatter)

	app := cli.NewApp()
	app.Name = "example"
	app.Usage = "a go module example"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		log.Notice("hello go-module")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Critical(err)
	}
}
