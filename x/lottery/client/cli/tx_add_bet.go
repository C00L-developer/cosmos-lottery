package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"lottery/x/lottery/types"
)

var _ = strconv.Itoa(0)

func CmdAddBet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-bet [amount] [suffix-hash]",
		Short: "Broadcast message add-bet",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argSuffixHash := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddBet(
				clientCtx.GetFromAddress().String(),
				argAmount,
				argSuffixHash,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
