package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"lottery/x/lottery/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddBet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-bet [amount] [suffix]",
		Short: "Broadcast message add-bet",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			suffixHash := sha256.Sum256([]byte(argAmount + args[1]))
			suffixHashString := hex.EncodeToString(suffixHash[:])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddBet(
				clientCtx.GetFromAddress().String(),
				argAmount,
				suffixHashString,
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
