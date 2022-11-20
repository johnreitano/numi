package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateAndVerifyUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-and-verify-user [user-id] [first-name] [last-name] [country-code] [subnational-entity] [city] [bio] [referrer] [account-address]",
		Short: "Broadcast message createAndVerifyUser",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUserId := args[0]
			argFirstName := args[1]
			argLastName := args[2]
			argCountryCode := args[3]
			argSubnationalEntity := args[4]
			argCity := args[5]
			argBio := args[6]
			argReferrer := args[7]
			argAccountAddress := args[8]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAndVerifyUser(
				clientCtx.GetFromAddress().String(),
				argUserId,
				argFirstName,
				argLastName,
				argCountryCode,
				argSubnationalEntity,
				argCity,
				argBio,
				argReferrer,
				argAccountAddress,
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
