package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/likecoin/likechain/x/likenft/types"
)

func (k msgServer) BuyNFT(goCtx context.Context, msg *types.MsgBuyNFT) (*types.MsgBuyNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	buyerAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf(err.Error())
	}

	// check listing exists
	sellerAddress, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return nil, types.ErrListingNotFound
	}
	listing, isFound := k.GetListing(ctx, msg.ClassId, msg.NftId, sellerAddress)
	if !isFound {
		return nil, types.ErrListingNotFound
	}

	// check listing owner is still valid
	if !k.nftKeeper.GetOwner(ctx, msg.ClassId, msg.NftId).Equals(sellerAddress) {
		return nil, types.ErrListingExpired.Wrapf("Listing owner is no longer valid")
	}

	// check listing not expired
	if listing.Expiration.Before(ctx.BlockTime()) {
		return nil, types.ErrListingExpired
	}

	// check price >= listing price
	if msg.Price < listing.Price {
		return nil, types.ErrFailedToBuyNFT.Wrapf("Price is too low. Listing price was %d", listing.Price)
	}

	// check user has enough balance
	if k.bankKeeper.GetBalance(ctx, buyerAddress, k.GetParams(ctx).MintPriceDenom).Amount.Uint64() < msg.Price {
		return nil, types.ErrFailedToBuyNFT.Wrapf("User does not have enough balance")
	}

	// transact
	// pay seller
	priceCoins := sdk.NewCoins(sdk.NewCoin(k.GetParams(ctx).MintPriceDenom, sdk.NewInt(int64(msg.Price))))
	err = k.bankKeeper.SendCoins(ctx, buyerAddress, sellerAddress, priceCoins)
	if err != nil {
		return nil, types.ErrFailedToBuyNFT.Wrapf(err.Error())
	}
	// TODO: pay royalty to class parent owner
	// transfer nft to buyer
	err = k.nftKeeper.Transfer(ctx, msg.ClassId, msg.NftId, buyerAddress)
	if err != nil {
		return nil, types.ErrFailedToBuyNFT.Wrapf(err.Error())
	}

	// remove listing
	k.RemoveListing(ctx, msg.ClassId, msg.NftId, sellerAddress)

	// owner changed, prune invalid listings
	k.PruneInvalidListingsForNFT(ctx, msg.ClassId, msg.NftId)

	// emit event
	ctx.EventManager().EmitTypedEvent(&types.EventBuyNFT{
		ClassId: msg.ClassId,
		NftId:   msg.NftId,
		Buyer:   buyerAddress.String(),
		Seller:  sellerAddress.String(),
		Price:   msg.Price,
	})

	return &types.MsgBuyNFTResponse{}, nil
}
