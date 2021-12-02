package cli

// DONTCOVER

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagPoolCoinDenom    = "pool-coin-denom"
	FlagReserveAcc       = "reserve-acc"
	FlagDepositorAddress = "depositor-address"
)

func flagSetPool() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagPoolCoinDenom, "", "The denomination of the pool coin")
	fs.String(FlagReserveAcc, "", "The Bech32 address of the reserve account")

	return fs
}

func flagSetDepositorAddress() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagDepositorAddress, "", "The Bech32 address of the depositor account")
	return fs
}
