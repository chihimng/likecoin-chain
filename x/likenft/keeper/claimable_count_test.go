package keeper_test

import (
	"fmt"
	"testing"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/golang/mock/gomock"
	"github.com/likecoin/likechain/backport/cosmos-sdk/v0.46.0-alpha2/x/nft"
	keepertest "github.com/likecoin/likechain/testutil/keeper"
	"github.com/likecoin/likechain/x/likenft/testutil"
	"github.com/likecoin/likechain/x/likenft/types"
)

type IUpdateClassClaimableCountMatcher struct {
	claimableCount uint64
}

func updateClassClaimableCountMatcher(claimableCount uint64) gomock.Matcher {
	return IUpdateClassClaimableCountMatcher{
		claimableCount,
	}
}

func (m IUpdateClassClaimableCountMatcher) Matches(x interface{}) bool {
	class := x.(nft.Class)
	var classData types.ClassData
	if err := classData.Unmarshal(class.Data.Value); err != nil {
		fmt.Printf(">> %s", err.Error())
		return false
	}
	fmt.Printf(">> %d", classData.ClaimableCount)
	return classData.ClaimableCount == m.claimableCount
}

func (m IUpdateClassClaimableCountMatcher) String() string {
	return fmt.Sprintf("data item claimableCount is equal to %d", m.claimableCount)
}

func getClassDataWithClaimableCount(count uint64) *cdctypes.Any {
	classData := types.ClassData{
		ClaimableCount: count,
	}
	classDataInAny, _ := cdctypes.NewAnyWithValue(&classData)
	return classDataInAny
}

func TestIncClaimableNFTCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	accountKeeper := testutil.NewMockAccountKeeper(ctrl)
	bankKeeper := testutil.NewMockBankKeeper(ctrl)
	iscnKeeper := testutil.NewMockIscnKeeper(ctrl)
	nftKeeper := testutil.NewMockNftKeeper(ctrl)
	keeper, ctx := keepertest.LikenftKeeperOverrideDependedKeepers(t, keepertest.LikenftDependedKeepers{
		AccountKeeper: accountKeeper,
		BankKeeper:    bankKeeper,
		IscnKeeper:    iscnKeeper,
		NftKeeper:     nftKeeper,
	})

	classId1 := "likenft1"
	claimableId1 := "claimable1"
	claimableId2 := "claimable2"

	gomock.InOrder(
		// create 1
		nftKeeper.
			EXPECT().
			GetClass(gomock.Any(), classId1).
			Return(nft.Class{
				Id:   classId1,
				Data: getClassDataWithClaimableCount(0),
			}, true),
		nftKeeper.
			EXPECT().
			UpdateClass(gomock.Any(), updateClassClaimableCountMatcher(1)).
			Return(nil),
		// update 1, no increment, no call to nft keeper
		// create 2
		nftKeeper.
			EXPECT().
			GetClass(gomock.Any(), classId1).
			Return(nft.Class{
				Id:   classId1,
				Data: getClassDataWithClaimableCount(1),
			}, true),
		nftKeeper.
			EXPECT().
			UpdateClass(gomock.Any(), updateClassClaimableCountMatcher(2)).
			Return(nil),
		// delete 2
		nftKeeper.
			EXPECT().
			GetClass(gomock.Any(), classId1).
			Return(nft.Class{
				Id:   classId1,
				Data: getClassDataWithClaimableCount(2),
			}, true),
		nftKeeper.
			EXPECT().
			UpdateClass(gomock.Any(), updateClassClaimableCountMatcher(1)),
		// delete 2 again, no decrement, no call to nft keeper
		// delete 1
		nftKeeper.
			EXPECT().
			GetClass(gomock.Any(), classId1).
			Return(nft.Class{
				Id:   classId1,
				Data: getClassDataWithClaimableCount(1),
			}, true),
		nftKeeper.
			EXPECT().
			UpdateClass(gomock.Any(), updateClassClaimableCountMatcher(0)),
	)

	// create 1
	keeper.SetClaimableNFT(ctx, types.ClaimableNFT{
		ClassId: classId1,
		Id:      claimableId1,
		Input:   types.NFTInput{},
	})
	// update 1
	keeper.SetClaimableNFT(ctx, types.ClaimableNFT{
		ClassId: classId1,
		Id:      claimableId1,
		Input:   types.NFTInput{},
	})
	// create 2
	keeper.SetClaimableNFT(ctx, types.ClaimableNFT{
		ClassId: classId1,
		Id:      claimableId2,
		Input:   types.NFTInput{},
	})
	// delete 2
	keeper.RemoveClaimableNFT(ctx, classId1, claimableId2)
	// delete 2 again
	keeper.RemoveClaimableNFT(ctx, classId1, claimableId2)
	// delete 1
	keeper.RemoveClaimableNFT(ctx, classId1, claimableId1)

	ctrl.Finish()
}
