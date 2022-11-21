package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/spf13/cobra"
)

func CmdListUserAccountAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-user-account-address",
		Short: "list all UserAccountAddress",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllUserAccountAddressRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.UserAccountAddressAll(context.Background(), params)
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

func CmdShowUserAccountAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-user-account-address [account-address]",
		Short: "shows a UserAccountAddress",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAccountAddress := args[0]

			params := &types.QueryGetUserAccountAddressRequest{
				AccountAddress: argAccountAddress,
			}

			res, err := queryClient.UserAccountAddress(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
