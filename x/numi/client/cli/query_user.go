package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/spf13/cobra"
)

func CmdListUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-user",
		Short: "list all User",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllUserRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.UserAll(context.Background(), params)
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

func CmdShowUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-user [user-id]",
		Short: "shows a User",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argUserId := args[0]

			params := &types.QueryGetUserRequest{
				UserId: argUserId,
			}

			res, err := queryClient.User(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
