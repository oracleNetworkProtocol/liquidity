package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	liquidity "github.com/oracleNetworkProtocol/liquidity/app"
	_ "github.com/oracleNetworkProtocol/liquidity/client/docs/statik"
	"github.com/oracleNetworkProtocol/liquidity/cmd/liquidityd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, liquidity.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
