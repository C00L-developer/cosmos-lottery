package cli

import (
	"context"
	"fmt"
	"strconv"

	"lottery/x/lottery/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLottery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-lottery",
		Short: "list all lottery",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLotteryRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LotteryAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowLottery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-lottery [index]",
		Short: "shows a lottery",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("lottery index %s not a valid uint, please input a valid lottery index", args[0])
			}

			params := &types.QueryGetLotteryRequest{
				Index: argIndex,
			}

			res, err := queryClient.Lottery(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
