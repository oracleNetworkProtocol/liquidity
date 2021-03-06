package cli

// DONTCOVER
// client is excluded from test coverage in the poc phase milestone 1 and will be included in milestone 2 with completeness

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/oracleNetworkProtocol/liquidity/x/liquidity/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	liquidityQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the liquidity module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	liquidityQueryCmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdQueryLiquidityPool(),
		GetCmdQueryLiquidityPools(),
		GetCmdQueryLiquidityPoolBatch(),
		GetCmdQueryPoolBatchDepositMsgs(),
		GetCmdQueryPoolBatchDepositMsg(),
		GetCmdQueryPoolBatchWithdrawMsgs(),
		GetCmdQueryPoolBatchWithdrawMsg(),
		GetCmdQueryPoolBatchSwapMsgs(),
		GetCmdQueryPoolBatchSwapMsg(),

		GetCmdQueryPoolDepositSuccessMsg(),
		GetCmdQueryPoolDepositSuccessMsgs(),
		GetCmdQueryPoolSwapSuccessMsg(),
		GetCmdQueryPoolSwapSuccessMsgs(),
		GetCmdQueryPoolWithdrawSuccessMsg(),
		GetCmdQueryPoolWithdrawSuccessMsgs(),
	)

	return liquidityQueryCmd
}

func GetCmdQueryPoolWithdrawSuccessMsgs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraws_finish [pool-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query all completed withdraw messages of liquidity pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all completed withdraws messages of liquidity pool on the specified pool

If finish messages are normally processed from the endblock, the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s withdraws_finish 1

Example (with withdraw address):
$ %[1]s query %[2]s withdraws_finish 1 --withdraw-address=[address]
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			withdrawAddress, _ := cmd.Flags().GetString(FlagWithdrawAddress)

			res, err := queryClient.PoolWithdrawSuccessMsgs(
				context.Background(),
				&types.QueryPoolWithdrawSuccessMsgsRequest{
					PoolId:          poolID,
					WithdrawAddress: withdrawAddress,
					Pagination:      pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	cmd.Flags().AddFlagSet(flagSetWithdrawAddress())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryPoolWithdrawSuccessMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw_finish [pool-id] [msg-index]",
		Args:  cobra.ExactArgs(2),
		Short: "Query all completed withdraw message of liquidity pool by msg index",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all completed withdraw message of liquidity pool on the specified pool by msg index

If finish messages are normally processed from the endblock, the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s withdraw_finish 1 1

`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint32, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			msgIndex, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("msg-index %s not a valid uint, input a valid unsigned 32-bit integer for msg-index", args[1])
			}

			res, err := queryClient.PoolWithdrawSuccessMsg(
				context.Background(),
				&types.QueryPoolWithdrawSuccessMsgRequest{
					PoolId:   poolID,
					MsgIndex: msgIndex,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryPoolSwapSuccessMsgs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swaps_finish [pool-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query all completed swap messages of liquidity pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all completed swap messages of liquidity pool on the specified pool

If finish messages are normally processed from the endblock, the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s swaps_finish 1

Example (with swap address):
$ %[1]s query %[2]s swaps_finish 1 --swap-address=[address]
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			swapAddress, _ := cmd.Flags().GetString(FlagSwapAddress)

			res, err := queryClient.PoolSwapSuccessMsgs(
				context.Background(),
				&types.QueryPoolSwapSuccessMsgsRequest{
					PoolId:      poolID,
					SwapAddress: swapAddress,
					Pagination:  pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	cmd.Flags().AddFlagSet(flagSetSwapAddress())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryPoolSwapSuccessMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap_finish [pool-id] [msg-index]",
		Args:  cobra.ExactArgs(2),
		Short: "Query all completed swap message of liquidity pool by msg index",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all completed swap message of liquidity pool on the specified pool by msg index

If finish messages are normally processed from the endblock, the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s swap_finish 1 1

`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint32, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			msgIndex, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("msg-index %s not a valid uint, input a valid unsigned 32-bit integer for msg-index", args[1])
			}

			res, err := queryClient.PoolSwapSuccessMsg(
				context.Background(),
				&types.QueryPoolSwapSuccessMsgRequest{
					PoolId:   poolID,
					MsgIndex: msgIndex,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryPoolDepositSuccessMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit_finish [pool-id] [msg-index]",
		Args:  cobra.ExactArgs(2),
		Short: "Query all completed deposit message of liquidity pool by msg index",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all completed deposit message of liquidity pool on the specified pool by msg index

If finish messages are normally processed from the endblock, the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s deposit_finish 1 1

`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint32, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			msgIndex, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("msg-index %s not a valid uint, input a valid unsigned 32-bit integer for msg-index", args[1])
			}

			res, err := queryClient.PoolDepositSuccessMsg(
				context.Background(),
				&types.QueryPoolDepositSuccessMsgRequest{
					PoolId:   poolID,
					MsgIndex: msgIndex,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryPoolDepositSuccessMsgs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposits_finish [pool-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query all completed deposit messages of liquidity pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all completed deposit messages of liquidity pool on the specified pool

If finish messages are normally processed from the endblock, the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s deposits_finish 1

Example (with depositor address):
$ %[1]s query %[2]s deposits_finish 1 --depositor-address=[address]
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			depositorAddress, _ := cmd.Flags().GetString(FlagDepositorAddress)

			res, err := queryClient.PoolDepositSuccessMsgs(
				context.Background(),
				&types.QueryPoolDepositSuccessMsgsRequest{
					PoolId:           poolID,
					DepositorAddress: depositorAddress,
					Pagination:       pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	cmd.Flags().AddFlagSet(flagSetDepositorAddress())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the values set as liquidity parameters",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as liquidity parameters.

Example:
$ %s query %s params
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(
				context.Background(),
				&types.QueryParamsRequest{},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryLiquidityPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool [pool-id]",
		Short: "Query details of a liquidity pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details of a liquidity pool
Example:
$ %[1]s query %[2]s pool 1

Example (with pool coin denom):
$ %[1]s query %[2]s pool --pool-coin-denom=[denom]

Example (with reserve acc):
$ %[1]s query %[2]s pool --reserve-acc=[address]
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			var res *types.QueryLiquidityPoolResponse
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			foundArg := false
			queryClient := types.NewQueryClient(clientCtx)

			poolCoinDenom, _ := cmd.Flags().GetString(FlagPoolCoinDenom)
			if poolCoinDenom != "" {
				foundArg = true
				res, err = queryClient.LiquidityPoolByPoolCoinDenom(
					context.Background(),
					&types.QueryLiquidityPoolByPoolCoinDenomRequest{PoolCoinDenom: poolCoinDenom},
				)
				if err != nil {
					return err
				}
			}

			reserveAcc, _ := cmd.Flags().GetString(FlagReserveAcc)
			if !foundArg && reserveAcc != "" {
				foundArg = true
				res, err = queryClient.LiquidityPoolByReserveAcc(
					context.Background(),
					&types.QueryLiquidityPoolByReserveAccRequest{ReserveAcc: reserveAcc},
				)
				if err != nil {
					return err
				}
			}

			if !foundArg && len(args) > 0 {
				poolID, err := strconv.ParseUint(args[0], 10, 64)
				if err != nil {
					return fmt.Errorf("pool-id %s not a valid uint, input a valid unsigned 32-bit integer for pool-id", args[0])
				}

				if poolID != 0 {
					foundArg = true
					res, err = queryClient.LiquidityPool(
						context.Background(),
						&types.QueryLiquidityPoolRequest{PoolId: poolID},
					)
					if err != nil {
						return err
					}
				}
			}

			if !foundArg {
				return fmt.Errorf("provide the pool-id argument or --%s or --%s flag", FlagPoolCoinDenom, FlagReserveAcc)
			}

			return clientCtx.PrintProto(res)
		},
	}
	cmd.Flags().AddFlagSet(flagSetPool())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryLiquidityPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pools",
		Args:  cobra.NoArgs,
		Short: "Query for all liquidity pools",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details about all liquidity pools on a network.
Example:
$ %s query %s pools
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.LiquidityPools(
				context.Background(),
				&types.QueryLiquidityPoolsRequest{Pagination: pageReq},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryLiquidityPoolBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch [pool-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query details of a liquidity pool batch",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details of a liquidity pool batch
Example:
$ %s query %s batch 1
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint32, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			res, err := queryClient.LiquidityPoolBatch(
				context.Background(),
				&types.QueryLiquidityPoolBatchRequest{PoolId: poolID},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdQueryPoolBatchDepositMsgs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposits [pool-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query all deposit messages of the liquidity pool batch",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all deposit messages of the liquidity pool batch on the specified pool

If batch messages are normally processed from the endblock, the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s deposits 1
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			res, err := queryClient.PoolBatchDepositMsgs(
				context.Background(),
				&types.QueryPoolBatchDepositMsgsRequest{
					PoolId:     poolID,
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryPoolBatchDepositMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [pool-id] [msg-index]",
		Args:  cobra.ExactArgs(2),
		Short: "Query the deposit messages on the liquidity pool batch",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the deposit messages on the liquidity pool batch for the specified pool-id and msg-index

If batch messages are normally processed from the endblock,
the resulting state is applied and the messages are removed from the beginning of the next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s deposit 1 20
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint, input a valid unsigned 32-bit integer for pool-id", args[0])
			}

			msgIndex, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("msg-index %s not a valid uint, input a valid unsigned 32-bit integer for msg-index", args[1])
			}

			res, err := queryClient.PoolBatchDepositMsg(
				context.Background(),
				&types.QueryPoolBatchDepositMsgRequest{
					PoolId:   poolID,
					MsgIndex: msgIndex,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryPoolBatchWithdrawMsgs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraws [pool-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for all withdraw messages on the liquidity pool batch",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all withdraw messages on the liquidity pool batch for the specified pool-id

If batch messages are normally processed from the endblock,
the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s withdraws 1
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			result, err := queryClient.PoolBatchWithdrawMsgs(context.Background(), &types.QueryPoolBatchWithdrawMsgsRequest{
				PoolId: poolID, Pagination: pageReq})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(result)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryPoolBatchWithdrawMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [pool-id] [msg-index]",
		Args:  cobra.ExactArgs(2),
		Short: "Query the withdraw messages in the liquidity pool batch",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the withdraw messages in the liquidity pool batch for the specified pool-id and msg-index

if the batch message are normally processed from the endblock,
the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s withdraw 1 20
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			msgIndex, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("msg-index %s not a valid uint, input a valid unsigned 32-bit integer msg-index", args[1])
			}

			res, err := queryClient.PoolBatchWithdrawMsg(
				context.Background(),
				&types.QueryPoolBatchWithdrawMsgRequest{
					PoolId:   poolID,
					MsgIndex: msgIndex,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryPoolBatchSwapMsgs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swaps [pool-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query all swap messages in the liquidity pool batch",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all swap messages in the liquidity pool batch for the specified pool-id

If batch messages are normally processed from the endblock,
the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s swaps 1
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint, input a valid unsigned 32-bit integer pool-id", args[0])
			}

			res, err := queryClient.PoolBatchSwapMsgs(
				context.Background(),
				&types.QueryPoolBatchSwapMsgsRequest{
					PoolId:     poolID,
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryPoolBatchSwapMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [pool-id] [msg-index]",
		Args:  cobra.ExactArgs(2),
		Short: "Query for the swap message on the batch of the liquidity pool specified pool-id and msg-index",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for the swap message on the batch of the liquidity pool specified pool-id and msg-index

If the batch message are normally processed and from the endblock,
the resulting state is applied and the messages are removed in the beginning of next block.
To query for past blocks, query the block height using the REST/gRPC API of a node that is not pruned.

Example:
$ %s query %s swap 1 20
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool-id %s not a valid uint, input a valid unsigned 32-bit integer for pool-id", args[0])
			}

			msgIndex, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("msg-index %s not a valid uint, input a valid unsigned 32-bit integer for msg-index", args[1])
			}

			res, err := queryClient.PoolBatchSwapMsg(
				context.Background(),
				&types.QueryPoolBatchSwapMsgRequest{
					PoolId:   poolID,
					MsgIndex: msgIndex,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
