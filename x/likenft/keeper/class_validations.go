package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/likecoin/likechain/x/likenft/types"
)

func (k Keeper) validateClassParentRelation(ctx sdk.Context, classId string, parent types.ClassParent) error {
	if parent.Type == types.ClassParentType_ISCN {
		classesByISCN, found := k.GetClassesByISCN(ctx, parent.IscnIdPrefix)
		if !found {
			return types.ErrNftClassNotRelatedToAnyIscn.Wrapf("NFT claims it is related to ISCN %s but no mapping is found", parent.IscnIdPrefix)
		}
		isRelated := false
		for _, validClassId := range classesByISCN.ClassIds {
			if validClassId == classId {
				// minted relation is valid
				isRelated = true
				break
			}
		}
		if !isRelated {
			return types.ErrNftClassNotRelatedToAnyIscn.Wrapf("NFT claims it is related to ISCN %s but no mapping is found", parent.IscnIdPrefix)
		}
	} else if parent.Type == types.ClassParentType_ACCOUNT {
		acc, err := sdk.AccAddressFromBech32(parent.Account)
		if err != nil {
			return sdkerrors.ErrInvalidAddress.Wrapf("%s", err.Error())
		}
		classesByAccount, found := k.GetClassesByAccount(ctx, acc)
		if !found {
			return types.ErrNftClassNotRelatedToAnyAccount.Wrapf("NFT claims it is related to account %s but no mapping is found", parent.Account)
		}
		isRelated := false
		for _, validClassId := range classesByAccount.ClassIds {
			if validClassId == classId {
				// minted relation is valid
				isRelated = true
				break
			}
		}
		if !isRelated {
			return types.ErrNftClassNotRelatedToAnyAccount.Wrapf("NFT claims it is related to account %s but no mapping is found", parent.Account)
		}
	} else {
		return sdkerrors.ErrInvalidRequest.Wrapf("Unsupported parent type %s in nft class", parent.Type.String())
	}
	return nil
}

func (k Keeper) resolveClassParentAndOwner(ctx sdk.Context, parentInput types.ClassParentInput, ownerBech32 string) (types.ClassParentAndOwner, error) {
	if parentInput.Type == types.ClassParentType_ISCN {
		iscnId, iscnRecord, err := k.resolveIscnIdAndRecord(ctx, parentInput.IscnIdPrefix)
		if err != nil {
			return types.ClassParentAndOwner{}, err
		}
		return types.ClassParentAndOwner{
			ClassParent: types.ClassParent{
				Type:              types.ClassParentType_ISCN,
				IscnIdPrefix:      iscnId.Prefix.String(),
				IscnVersionAtMint: iscnRecord.LatestVersion,
			},
			Owner: iscnRecord.OwnerAddress(),
		}, nil
	} else if parentInput.Type == types.ClassParentType_ACCOUNT {
		owner, err := sdk.AccAddressFromBech32(ownerBech32)
		if err != nil {
			return types.ClassParentAndOwner{}, sdkerrors.ErrInvalidAddress.Wrapf("%s", err.Error())
		}
		return types.ClassParentAndOwner{
			ClassParent: types.ClassParent{
				Type:    types.ClassParentType_ACCOUNT,
				Account: owner.String(),
			},
			Owner: owner,
		}, nil
	} else {
		return types.ClassParentAndOwner{}, sdkerrors.ErrInvalidRequest.Wrapf("Unsupported parent type %s in nft class", parentInput.Type.String())
	}
}

func (k Keeper) validateAndGetClassParentAndOwner(ctx sdk.Context, classId string, classData *types.ClassData) (*types.ClassParentAndOwner, error) {
	if err := k.validateClassParentRelation(ctx, classId, classData.Parent); err != nil {
		return nil, err
	}

	// refresh parent info (e.g. iscn latest version) & check ownership
	parent, err := k.resolveClassParentAndOwner(ctx, classData.Parent.ToInput(), classData.Parent.Account)
	if err != nil {
		return nil, err
	}
	return &parent, nil
}

func (k msgServer) sanitizeBlindBoxConfig(blindBoxConfig *types.BlindBoxConfig) (*types.BlindBoxConfig, error) {
	if blindBoxConfig == nil {
		return nil, nil
	}
	if len(blindBoxConfig.MintPeriods) <= 0 {
		return nil, types.ErrInvalidNftClassConfig.Wrapf("Mint period cannot be empty")
	}
	// Sort the mint period by start time
	blindBoxConfig.MintPeriods = SortMintPeriod(blindBoxConfig.MintPeriods, true)
	for _, mintPeriod := range blindBoxConfig.MintPeriods {
		// Ensure all mint period start time is before reveal time
		if mintPeriod.StartTime.After(blindBoxConfig.RevealTime) {
			return nil, types.ErrInvalidNftClassConfig.Wrapf("One of the mint periods' start time %s is after reveal time %s", mintPeriod.StartTime.String(), blindBoxConfig.RevealTime.String())
		}
		// Ensure all the addresses in allow list is valid
		for _, allowedAddress := range mintPeriod.AllowedAddresses {
			if _, err := sdk.AccAddressFromBech32(allowedAddress); err != nil {
				return nil, sdkerrors.ErrInvalidAddress.Wrapf("One of the allowed addresses %s is invalid", allowedAddress)
			}
		}
	}
	return blindBoxConfig, nil
}

func (k msgServer) sanitizeClassConfig(classConfig types.ClassConfig) (*types.ClassConfig, error) {
	// Ensure mint periods and reveal time are set when blind box mode is enabled
	cleanBlindBoxConfig, err := k.sanitizeBlindBoxConfig(classConfig.BlindBoxConfig)
	if err != nil {
		return nil, err
	}
	classConfig.BlindBoxConfig = cleanBlindBoxConfig
	return &classConfig, nil
}
