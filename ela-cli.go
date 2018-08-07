package main

import (
	"os"
	"sort"

	"github.com/elastos/Elastos.ELA.Client/cli/info"
	"github.com/elastos/Elastos.ELA.Client/cli/wallet"
	"github.com/elastos/Elastos.ELA.Client/cli/mine"
	"github.com/elastos/Elastos.ELA.Client/cli/smartcontract"
	"github.com/elastos/Elastos.ELA.Client/log"
	cliLog "github.com/elastos/Elastos.ELA.Client/cli/log"
	"github.com/urfave/cli"
)

var Version string

func init() {
	log.InitLog()
}

func main() {
	app := cli.NewApp()
	app.Name = "ela-cli"
	app.Version = Version
	app.HelpName = "ela-cli"
	app.Usage = "command line tool for ELA blockchain"
	app.UsageText = "ela-cli [global options] command [command options] [args]"
	app.HideHelp = false
	app.HideVersion = false
	//commands
	app.Commands = []cli.Command{
		*cliLog.NewCommand(),
		*info.NewCommand(),
		*wallet.NewCommand(),
		*mine.NewCommand(),
		*smartcontract.NewCommand(),
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}
