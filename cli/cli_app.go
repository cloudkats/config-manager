package cli

import (
	"io"

	"github.com/urfave/cli/v2"
)

func CreateCli(version string, writer io.Writer, errwriter io.Writer) *cli.App {
	cli.OsExiter = func(exitCode int) {
		// Do nothing. We just need to override this function, as the default value calls os.Exit, which
		// kills the app (or any automated test) dead in its tracks.
	}

	cli.AppHelpTemplate = "contrive - demonstrating the available API"

	app := cli.NewApp()

	app.Name = "config-manager"
	app.Authors = []*cli.Author{
		{
			Name:  "Cloud Kats Labs",
			Email: "cloudkats@gmail.com",
		},
	}
	app.Copyright = "(c) 2021 Cloud Kats Labs"
	app.Version = version
	app.EnableBashCompletion = true
	app.Usage = "config-manager <COMMAND> [GLOBAL OPTIONS]"
	app.Writer = writer
	app.ErrWriter = errwriter
	app.UsageText = `For documentation, see https://github.com/cloudkats/config-manager.`
	app.Commands = configFunctions()

	return app
}
