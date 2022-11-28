package mock_types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/golang/mock/gomock"
)

func (bankKeeper *MockBankKeeper) ExpectAny(context context.Context) {
	bankKeeper.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	bankKeeper.EXPECT().MintCoins(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any()).AnyTimes()
}

func (mintKeeper *MockMintKeeper) ExpectAny(context context.Context) {
	mintKeeper.EXPECT().GetParams(sdk.UnwrapSDKContext(context)).Return(minttypes.Params{MintDenom: "uatom"}).AnyTimes()
}
