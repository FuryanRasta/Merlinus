package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"

	"github.com/magewar/mage/app"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/magewar/mage/app/mage/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
