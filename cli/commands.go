package cli

import (
	"errors"

	"config-manager/internal"
	"config-manager/internal/commands"
	"github.com/urfave/cli/v2"
)

// TODO
// test
func configFunctions() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "context",
			Description: "Read context in directory, merge values and provide a visualisation for given json path",
			HelpName:    "config-manager context --context test/fixtures/context/users",
			ArgsUsage:   "<json path>",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "context",
					Aliases:  []string{"c"},
					Required: true,
					Usage:    "Location where to look for context"},
			},
			Action: func(ctx *cli.Context) error {
				if ctx.Args().Len() > 1 {
					return errors.New("requires at max single argument: cli context <key>")
				}

				ctxPath := ctx.String("context")
				key := ctx.Args().First()
				// fmt.Printf("Json Path %q\n", key)
				// fmt.Printf("Context %q\n", ctxPath)

				scope := internal.NewScope()
				cmd := commands.New()

				if ctx.Args().Len() == 1 {
					return cmd.ContextLookup(scope, ctxPath, key)
				} else {
					return cmd.FullContext(scope, ctxPath)
				}
			},
		},
	}
}
