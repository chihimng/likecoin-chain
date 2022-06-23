package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/likecoin/likechain/x/likenft/types"
)

func (k msgServer) validateReqToMutateRoyaltyConfig(ctx sdk.Context, msgCreator string, classId string) error {
	// Check user is class owner
	userAddress, err := sdk.AccAddressFromBech32(msgCreator)
	if err != nil {
		return sdkerrors.ErrInvalidAddress
	}
	class, classData, err := k.GetClass(ctx, classId)
	if err != nil {
		return err
	}
	parent, err := k.ValidateAndRefreshClassParent(ctx, class.Id, classData.Parent)
	if err != nil {
		return err
	}
	if !parent.Owner.Equals(userAddress) {
		return sdkerrors.ErrUnauthorized.Wrapf("user is not the class owner")
	}
	return nil
}

func (k msgServer) CreateRoyaltyConfig(goCtx context.Context, msg *types.MsgCreateRoyaltyConfig) (*types.MsgCreateRoyaltyConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.validateReqToMutateRoyaltyConfig(ctx, msg.Creator, msg.ClassId)
	if err != nil {
		return nil, err
	}

	// Check if the value already exists
	_, isFound := k.GetRoyaltyConfig(
		ctx,
		msg.ClassId,
	)
	if isFound {
		return nil, types.ErrRoyaltyConfigAlreadyExists
	}

	// Set config
	config := msg.RoyaltyConfig.ToConfig()
	k.SetRoyaltyConfig(
		ctx,
		types.RoyaltyConfigByClass{
			ClassId:       msg.ClassId,
			RoyaltyConfig: config,
		},
	)

	// TODO emit event

	return &types.MsgCreateRoyaltyConfigResponse{
		RoyaltyConfig: config,
	}, nil
}

func (k msgServer) UpdateRoyaltyConfig(goCtx context.Context, msg *types.MsgUpdateRoyaltyConfig) (*types.MsgUpdateRoyaltyConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.validateReqToMutateRoyaltyConfig(ctx, msg.Creator, msg.ClassId)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	_, isFound := k.GetRoyaltyConfig(
		ctx,
		msg.ClassId,
	)
	if !isFound {
		return nil, types.ErrRoyaltyConfigNotFound
	}

	config := msg.RoyaltyConfig.ToConfig()
	k.SetRoyaltyConfig(
		ctx,
		types.RoyaltyConfigByClass{
			ClassId:       msg.ClassId,
			RoyaltyConfig: config,
		},
	)

	// TODO emit event

	return &types.MsgUpdateRoyaltyConfigResponse{
		RoyaltyConfig: config,
	}, nil
}

func (k msgServer) DeleteRoyaltyConfig(goCtx context.Context, msg *types.MsgDeleteRoyaltyConfig) (*types.MsgDeleteRoyaltyConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.validateReqToMutateRoyaltyConfig(ctx, msg.Creator, msg.ClassId)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	_, isFound := k.GetRoyaltyConfig(
		ctx,
		msg.ClassId,
	)
	if !isFound {
		return nil, types.ErrRoyaltyConfigNotFound
	}

	k.RemoveRoyaltyConfig(
		ctx,
		msg.ClassId,
	)

	// TODO emit event

	return &types.MsgDeleteRoyaltyConfigResponse{}, nil
}
