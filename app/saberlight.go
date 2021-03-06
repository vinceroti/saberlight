package main

import (
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/commands"
)

func main() {
	switch kingpin.MustParse(cli.App.Parse(os.Args[1:])) {
	case cli.Version.FullCommand():
		commands.Version()
	case cli.Scan.FullCommand():
		commands.Scan()
	case cli.On.FullCommand():
		commands.On()
	case cli.Off.FullCommand():
		commands.Off()
	case cli.Dump.FullCommand():
		commands.Dump()
	case cli.Status.FullCommand():
		commands.Status()
	case cli.Color.FullCommand():
		commands.Color()
	case cli.White.FullCommand():
		commands.White()
	case cli.Mode.FullCommand():
		commands.Mode()
	case cli.Time.FullCommand():
		commands.Time()
	case cli.Schedule.FullCommand():
		commands.Schedule()
	}
}
