package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/likecoin/likechain/x/likenft/types"
)

func (k msgServer) DeleteMintableNFT(goCtx context.Context, msg *types.MsgDeleteMintableNFT) (*types.MsgDeleteMintableNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	class, classData, err := k.GetClass(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}
	parent, err := k.ValidateAndRefreshClassParent(ctx, msg.ClassId, classData.Parent)
	if err != nil {
		return nil, err
	}
	if err := k.validateReqToMutateMintableNFT(ctx, msg.Creator, class, classData, parent, false); err != nil {
		return nil, err
	}

	// check id already exists
	if _, exists := k.GetMintableNFT(ctx, msg.ClassId, msg.Id); !exists {
		return nil, types.ErrMintableNftNotFound
	}

	// remove record
	k.RemoveMintableNFT(ctx, msg.ClassId, msg.Id)

	// Emit event
	ctx.EventManager().EmitTypedEvent(&types.EventDeleteMintableNFT{
		ClassId:                 msg.ClassId,
		MintableNftId:           msg.Id,
		ClassParentIscnIdPrefix: parent.IscnIdPrefix,
		ClassParentAccount:      parent.Account,
	})

	return &types.MsgDeleteMintableNFTResponse{}, nil
}