package main

import (
	"fmt"
	"os"

	"config-manager/cli"
)

func main() {

	app := cli.CreateCli("0.0.1", os.Stdout, os.Stderr)
	err := app.Run(os.Args)

	checkForErrorsAndExit(err)
}

// If there is an error, display it in the console and exit with a non-zero exit code. Otherwise, exit 0.
func checkForErrorsAndExit(err error) {
	if err == nil {
		os.Exit(0)
	} else {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
