package keeper_test

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"time"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/likecoin/likechain/backport/cosmos-sdk/v0.46.0-alpha2/x/nft"
	apptestutil "github.com/likecoin/likechain/testutil"
	keepertest "github.com/likecoin/likechain/testutil/keeper"
	"github.com/likecoin/likechain/x/likenft"
	"github.com/likecoin/likechain/x/likenft/testutil"
	"github.com/likecoin/likechain/x/likenft/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// Test feature

func parseRevealEvent(t *testing.T, ctx sdk.Context) *types.EventRevealClass {
	for _, event := range ctx.EventManager().Events() {
		if event.Type != "likechain.likenft.EventRevealClass" {
			continue
		}
		ev := types.EventRevealClass{}
		for _, attr := range event.Attributes {
			val := strings.Trim(string(attr.Value), "\"")
			if string(attr.Key) == "class_id" {
				ev.ClassId = val
				continue
			}
			if string(attr.Key) == "success" {
				ev.Success = val == "true"
				continue
			}
			if string(attr.Key) == "error" {
				ev.Error = val
				continue
			}
		}
		return &ev
	}
	return nil
}

func TestRevealFeature(t *testing.T) {
	app := apptestutil.SetupTestAppWithDefaultState()
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	header := tmproto.Header{Height: app.LastBlockHeight() + 1}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	// seed class, relation
	ownerAddressBytes := []byte{0, 1, 0, 1, 0, 1, 0, 1}
	ownerAddress, _ := sdk.Bech32ifyAddressBytes("like", ownerAddressBytes)
	classId := "likenft11"
	classParent := types.ClassParent{
		Type:    types.ClassParentType_ACCOUNT,
		Account: ownerAddress,
	}
	classData := types.ClassData{
		Metadata: types.JsonInput(`{"abc":123}`),
		Parent:   classParent,
		Config: types.ClassConfig{
			Burnable:  true,
			MaxSupply: 3,
			BlindBoxConfig: &types.BlindBoxConfig{
				MintPeriods: []types.MintPeriod{
					{
						StartTime:        time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
						AllowedAddresses: []string{},
						MintPrice:        0,
					},
				},
				RevealTime: time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		MintableCount: 0, // will be incremented on seed mintables
		ToBeRevealed:  true,
	}
	classDataInAny, err := cdctypes.NewAnyWithValue(&classData)
	require.NoError(t, err)
	app.NftKeeper.SaveClass(ctx, nft.Class{
		Id:          classId,
		Name:        "Test Class",
		Symbol:      "TEST",
		Description: "Class for testing",
		Uri:         "ipfs://123456",
		UriHash:     "123456",
		Data:        classDataInAny,
	})
	app.LikeNftKeeper.SetClassesByAccount(ctx, types.ClassesByAccount{
		Account:  ownerAddress,
		ClassIds: []string{classId},
	})
	app.LikeNftKeeper.SetClassRevealQueueEntry(ctx, types.ClassRevealQueueEntry{
		RevealTime: time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
		ClassId:    classId,
	})

	// seed mintables

	app.LikeNftKeeper.SetMintableNFT(ctx, types.MintableNFT{
		ClassId: classId,
		Id:      "mintable1",
		Input: types.NFTInput{
			Uri:      "https://testnft.com/1",
			UriHash:  "1",
			Metadata: types.JsonInput(`1`),
		},
	})
	app.LikeNftKeeper.SetMintableNFT(ctx, types.MintableNFT{
		ClassId: classId,
		Id:      "mintable2",
		Input: types.NFTInput{
			Uri:      "https://testnft.com/2",
			UriHash:  "2",
			Metadata: types.JsonInput(`2`),
		},
	})
	app.LikeNftKeeper.SetMintableNFT(ctx, types.MintableNFT{
		ClassId: classId,
		Id:      "mintable3",
		Input: types.NFTInput{
			Uri:      "https://testnft.com/3",
			UriHash:  "3",
			Metadata: types.JsonInput(`3`),
		},
	})
	app.LikeNftKeeper.SetMintableNFT(ctx, types.MintableNFT{
		ClassId: classId,
		Id:      "mintable4",
		Input: types.NFTInput{
			Uri:      "https://testnft.com/4",
			UriHash:  "4",
			Metadata: types.JsonInput(`4`),
		},
	})
	app.LikeNftKeeper.SetMintableNFT(ctx, types.MintableNFT{
		ClassId: classId,
		Id:      "mintable5",
		Input: types.NFTInput{
			Uri:      "https://testnft.com/5",
			UriHash:  "5",
			Metadata: types.JsonInput(`5`),
		},
	})

	// seed tokens minted by buyers
	userAddressBytes := []byte{1, 1, 1, 1, 1, 1, 1, 1}

	nftData := types.NFTData{
		ClassParent:  classParent,
		ToBeRevealed: true,
	}
	nftDataInAny, err := cdctypes.NewAnyWithValue(&nftData)
	require.NoError(t, err)

	err = app.NftKeeper.Mint(ctx, nft.NFT{
		ClassId: classId,
		Id:      "nft1",
		Data:    nftDataInAny,
	}, userAddressBytes)
	require.NoError(t, err)

	err = app.NftKeeper.Mint(ctx, nft.NFT{
		ClassId: classId,
		Id:      "nft2",
		Data:    nftDataInAny,
	}, userAddressBytes)
	require.NoError(t, err)

	require.Equal(t, uint64(2), app.NftKeeper.GetTotalSupply(ctx, classId))

	// Increase height to after reveal time
	newHeader := ctx.BlockHeader()
	// hash from mainnet block 1
	hash, err := hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")
	require.NoError(t, err)
	newHeader.AppHash = hash
	newHeader.Time = time.Date(2022, 2, 1, 0, 0, 1, 0, time.UTC)
	ctx = ctx.WithBlockHeader(newHeader)

	require.NotPanics(t, func() {
		likenft.EndBlocker(ctx, app.LikeNftKeeper)
	})

	// check result from event
	event := parseRevealEvent(t, ctx)
	require.NotNil(t, event)
	require.Empty(t, event.Error)
	require.Equal(t, classId, event.ClassId)
	require.True(t, event.Success)

	// check nft content
	uriSeq := []string{}
	uriHashSeq := []string{}
	metadataSeq := []types.JsonInput{}

	for i := 0; i < 5; i++ {
		nft, found := app.NftKeeper.GetNFT(ctx, classId, fmt.Sprintf("nft%d", i+1))
		require.True(t, found)
		require.NotEmpty(t, nft.Uri)
		require.NotEmpty(t, nft.UriHash)
		uriSeq = append(uriSeq, nft.Uri)
		uriHashSeq = append(uriHashSeq, nft.UriHash)
		var nftData types.NFTData
		err := nftData.Unmarshal(nft.Data.Value)
		require.NoError(t, err)
		require.Equal(t, classParent, nftData.ClassParent)
		require.NotEmpty(t, nftData.Metadata)
		require.False(t, nftData.ToBeRevealed)
		metadataSeq = append(metadataSeq, nftData.Metadata)
	}

	// check sequence is randomized
	require.NotEqual(t, []string{"https://testnft.com/1", "https://testnft.com/2", "https://testnft.com/3", "https://testnft.com/4", "https://testnft.com/5"}, uriSeq)
	require.NotEqual(t, []string{"1", "2", "3", "4", "5"}, uriHashSeq)
	require.NotEqual(t, []types.JsonInput{types.JsonInput(`1`), types.JsonInput(`2`), types.JsonInput(`3`), types.JsonInput(`4`), types.JsonInput(`5`)}, metadataSeq)

	// check mintables removed
	mintables := app.LikeNftKeeper.GetMintableNFTs(ctx, classId)
	require.Empty(t, mintables)
}

// Tests control flow & external call counts

func TestRevealNormalMintToOwner(t *testing.T) {
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
	// hash from mainnet block 1
	hash, err := hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")
	require.NoError(t, err)
	ctx = ctx.WithBlockHeader(tmproto.Header{
		AppHash: hash,
	})

	ownerAddressBytes := []byte{0, 1, 0, 1, 0, 1, 0, 1}
	ownerAddress, _ := sdk.Bech32ifyAddressBytes("cosmos", ownerAddressBytes)
	classId := "likenft1aabbccddeeff"
	supply := 100
	mintableCount := 99
	totalSupply := 90
	mintToOwnerCount := mintableCount - totalSupply

	classData := types.ClassData{
		Metadata: types.JsonInput(`{"aaaa": "bbbb"}`),
		Parent: types.ClassParent{
			Type:    types.ClassParentType_ACCOUNT,
			Account: ownerAddress,
		},
		Config: types.ClassConfig{
			Burnable:  false,
			MaxSupply: uint64(supply),
			BlindBoxConfig: &types.BlindBoxConfig{
				MintPeriods: []types.MintPeriod{
					{
						StartTime:        time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
						AllowedAddresses: []string{},
						MintPrice:        uint64(0),
					},
				},
			},
		},
		MintableCount: uint64(mintableCount),
		ToBeRevealed:  true,
	}
	classDataInAny, _ := cdctypes.NewAnyWithValue(&classData)
	class := nft.Class{
		Id:          classId,
		Name:        "Class Name",
		Symbol:      "ABC",
		Description: "Testing Class 123",
		Uri:         "ipfs://abcdef",
		UriHash:     "abcdef",
		Data:        classDataInAny,
	}
	// mock calls
	nftKeeper.EXPECT().GetClass(gomock.Any(), classId).Return(class, true).AnyTimes()
	keeper.SetClassesByAccount(ctx, types.ClassesByAccount{
		Account:  ownerAddress,
		ClassIds: []string{classId},
	})
	nftKeeper.EXPECT().GetTotalSupply(ctx, classId).Return(uint64(totalSupply))
	nftKeeper.EXPECT().Mint(gomock.Any(), gomock.Any(), ownerAddressBytes).Return(nil).Times(mintToOwnerCount)
	nftKeeper.EXPECT().UpdateClass(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	for i := 0; i < mintableCount; i++ {
		keeper.SetMintableNFT(ctx, types.MintableNFT{
			ClassId: classId,
			Id:      fmt.Sprintf("mintable%d", i),
			Input: types.NFTInput{
				Uri: fmt.Sprintf("mintable%d", i),
			},
		})
	}
	var blindTokens []nft.NFT
	for i := 0; i < totalSupply+mintToOwnerCount; i++ {
		nftData := types.NFTData{
			ClassParent:  classData.Parent,
			ToBeRevealed: true,
		}
		nftDataInAny, err := cdctypes.NewAnyWithValue(&nftData)
		require.NoError(t, err)
		blindTokens = append(blindTokens, nft.NFT{
			ClassId: classId,
			Id:      fmt.Sprintf("nft%d", i),
			Data:    nftDataInAny,
		})
	}
	nftKeeper.EXPECT().GetNFTsOfClass(ctx, classId).Return(blindTokens)
	nftKeeper.EXPECT().Update(ctx, gomock.Any()).Return(nil).Times(totalSupply + mintToOwnerCount)

	// call
	err = keeper.RevealMintableNFTs(ctx, classId)
	require.NoError(t, err)

	ctrl.Finish()
}

func TestRevealNormalNoMintToOwner(t *testing.T) {
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
	// hash from mainnet block 1
	hash, err := hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")
	require.NoError(t, err)
	ctx = ctx.WithBlockHeader(tmproto.Header{
		AppHash: hash,
	})

	ownerAddressBytes := []byte{0, 1, 0, 1, 0, 1, 0, 1}
	ownerAddress, _ := sdk.Bech32ifyAddressBytes("cosmos", ownerAddressBytes)
	classId := "likenft1aabbccddeeff"
	supply := 100
	mintableCount := 99
	totalSupply := 99

	classData := types.ClassData{
		Metadata: types.JsonInput(`{"aaaa": "bbbb"}`),
		Parent: types.ClassParent{
			Type:    types.ClassParentType_ACCOUNT,
			Account: ownerAddress,
		},
		Config: types.ClassConfig{
			Burnable:  false,
			MaxSupply: uint64(supply),
			BlindBoxConfig: &types.BlindBoxConfig{
				MintPeriods: []types.MintPeriod{
					{
						StartTime:        time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
						AllowedAddresses: []string{},
						MintPrice:        uint64(0),
					},
				},
			},
		},
		MintableCount: uint64(mintableCount),
		ToBeRevealed:  true,
	}
	classDataInAny, _ := cdctypes.NewAnyWithValue(&classData)
	class := nft.Class{
		Id:          classId,
		Name:        "Class Name",
		Symbol:      "ABC",
		Description: "Testing Class 123",
		Uri:         "ipfs://abcdef",
		UriHash:     "abcdef",
		Data:        classDataInAny,
	}
	// mock calls
	nftKeeper.EXPECT().GetClass(gomock.Any(), classId).Return(class, true).AnyTimes()
	keeper.SetClassesByAccount(ctx, types.ClassesByAccount{
		Account:  ownerAddress,
		ClassIds: []string{classId},
	})
	nftKeeper.EXPECT().GetTotalSupply(ctx, classId).Return(uint64(totalSupply))
	nftKeeper.EXPECT().UpdateClass(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	for i := 0; i < mintableCount; i++ {
		keeper.SetMintableNFT(ctx, types.MintableNFT{
			ClassId: classId,
			Id:      fmt.Sprintf("mintable%d", i),
			Input: types.NFTInput{
				Uri: fmt.Sprintf("mintable%d", i),
			},
		})
	}
	var blindTokens []nft.NFT
	for i := 0; i < totalSupply; i++ {
		nftData := types.NFTData{
			ClassParent:  classData.Parent,
			ToBeRevealed: true,
		}
		nftDataInAny, err := cdctypes.NewAnyWithValue(&nftData)
		require.NoError(t, err)
		blindTokens = append(blindTokens, nft.NFT{
			ClassId: classId,
			Id:      fmt.Sprintf("nft%d", i),
			Data:    nftDataInAny,
		})
	}
	nftKeeper.EXPECT().GetNFTsOfClass(ctx, classId).Return(blindTokens)
	nftKeeper.EXPECT().Update(ctx, gomock.Any()).Return(nil).Times(totalSupply)

	// call
	err = keeper.RevealMintableNFTs(ctx, classId)
	require.NoError(t, err)

	ctrl.Finish()
}

func TestRevealClassNotFound(t *testing.T) {
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
	// hash from mainnet block 1
	hash, err := hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")
	require.NoError(t, err)
	ctx = ctx.WithBlockHeader(tmproto.Header{
		AppHash: hash,
	})

	classId := "likenft1aabbccddeeff"

	// mock calls
	nftKeeper.EXPECT().GetClass(gomock.Any(), classId).Return(nft.Class{}, false)

	// call
	err = keeper.RevealMintableNFTs(ctx, classId)
	require.Error(t, err)
	require.Contains(t, err.Error(), types.ErrNftClassNotFound.Error())

	ctrl.Finish()
}

func TestRevealNotBlindBox(t *testing.T) {
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
	// hash from mainnet block 1
	hash, err := hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")
	require.NoError(t, err)
	ctx = ctx.WithBlockHeader(tmproto.Header{
		AppHash: hash,
	})

	ownerAddressBytes := []byte{0, 1, 0, 1, 0, 1, 0, 1}
	ownerAddress, _ := sdk.Bech32ifyAddressBytes("cosmos", ownerAddressBytes)
	classId := "likenft1aabbccddeeff"
	supply := 100

	classData := types.ClassData{
		Metadata: types.JsonInput(`{"aaaa": "bbbb"}`),
		Parent: types.ClassParent{
			Type:    types.ClassParentType_ACCOUNT,
			Account: ownerAddress,
		},
		Config: types.ClassConfig{
			Burnable:       false,
			MaxSupply:      uint64(supply),
			BlindBoxConfig: nil,
		},
		MintableCount: uint64(0),
		ToBeRevealed:  false,
	}
	classDataInAny, _ := cdctypes.NewAnyWithValue(&classData)
	class := nft.Class{
		Id:          classId,
		Name:        "Class Name",
		Symbol:      "ABC",
		Description: "Testing Class 123",
		Uri:         "ipfs://abcdef",
		UriHash:     "abcdef",
		Data:        classDataInAny,
	}
	// mock calls
	nftKeeper.EXPECT().GetClass(gomock.Any(), classId).Return(class, true).AnyTimes()
	keeper.SetClassesByAccount(ctx, types.ClassesByAccount{
		Account:  ownerAddress,
		ClassIds: []string{classId},
	})

	// call
	err = keeper.RevealMintableNFTs(ctx, classId)
	require.Error(t, err)
	require.Contains(t, err.Error(), types.ErrClassIsNotBlindBox.Error())

	ctrl.Finish()
}

// Note: validateAndGetClassParentAndOwner covered by other cases
// FIXME: refactor to test that utils separately

func TestRevealFailedToMint(t *testing.T) {
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
	// hash from mainnet block 1
	hash, err := hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")
	require.NoError(t, err)
	ctx = ctx.WithBlockHeader(tmproto.Header{
		AppHash: hash,
	})

	ownerAddressBytes := []byte{0, 1, 0, 1, 0, 1, 0, 1}
	ownerAddress, _ := sdk.Bech32ifyAddressBytes("cosmos", ownerAddressBytes)
	classId := "likenft1aabbccddeeff"
	supply := 100
	mintableCount := 99
	totalSupply := 90

	classData := types.ClassData{
		Metadata: types.JsonInput(`{"aaaa": "bbbb"}`),
		Parent: types.ClassParent{
			Type:    types.ClassParentType_ACCOUNT,
			Account: ownerAddress,
		},
		Config: types.ClassConfig{
			Burnable:  false,
			MaxSupply: uint64(supply),
			BlindBoxConfig: &types.BlindBoxConfig{
				MintPeriods: []types.MintPeriod{
					{
						StartTime:        time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
						AllowedAddresses: []string{},
						MintPrice:        uint64(0),
					},
				},
			},
		},
		MintableCount: uint64(mintableCount),
		ToBeRevealed:  true,
	}
	classDataInAny, _ := cdctypes.NewAnyWithValue(&classData)
	class := nft.Class{
		Id:          classId,
		Name:        "Class Name",
		Symbol:      "ABC",
		Description: "Testing Class 123",
		Uri:         "ipfs://abcdef",
		UriHash:     "abcdef",
		Data:        classDataInAny,
	}
	// mock calls
	nftKeeper.EXPECT().GetClass(gomock.Any(), classId).Return(class, true).AnyTimes()
	keeper.SetClassesByAccount(ctx, types.ClassesByAccount{
		Account:  ownerAddress,
		ClassIds: []string{classId},
	})
	nftKeeper.EXPECT().GetTotalSupply(ctx, classId).Return(uint64(totalSupply))
	nftKeeper.EXPECT().Mint(gomock.Any(), gomock.Any(), ownerAddressBytes).Return(fmt.Errorf("Failed to mint"))

	// call
	err = keeper.RevealMintableNFTs(ctx, classId)
	require.Error(t, err)
	require.Contains(t, err.Error(), types.ErrFailedToMintNFT.Error())

	ctrl.Finish()
}

func TestRevealMintableMismatch(t *testing.T) {
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
	// hash from mainnet block 1
	hash, err := hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")
	require.NoError(t, err)
	ctx = ctx.WithBlockHeader(tmproto.Header{
		AppHash: hash,
	})

	ownerAddressBytes := []byte{0, 1, 0, 1, 0, 1, 0, 1}
	ownerAddress, _ := sdk.Bech32ifyAddressBytes("cosmos", ownerAddressBytes)
	classId := "likenft1aabbccddeeff"
	supply := 100
	mintableCount := 99
	totalSupply := 90
	mintToOwnerCount := mintableCount - totalSupply

	classData := types.ClassData{
		Metadata: types.JsonInput(`{"aaaa": "bbbb"}`),
		Parent: types.ClassParent{
			Type:    types.ClassParentType_ACCOUNT,
			Account: ownerAddress,
		},
		Config: types.ClassConfig{
			Burnable:  false,
			MaxSupply: uint64(supply),
			BlindBoxConfig: &types.BlindBoxConfig{
				MintPeriods: []types.MintPeriod{
					{
						StartTime:        time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
						AllowedAddresses: []string{},
						MintPrice:        uint64(0),
					},
				},
			},
		},
		MintableCount: uint64(mintableCount),
		ToBeRevealed:  true,
	}
	classDataInAny, _ := cdctypes.NewAnyWithValue(&classData)
	class := nft.Class{
		Id:          classId,
		Name:        "Class Name",
		Symbol:      "ABC",
		Description: "Testing Class 123",
		Uri:         "ipfs://abcdef",
		UriHash:     "abcdef",
		Data:        classDataInAny,
	}
	// mock calls
	nftKeeper.EXPECT().GetClass(gomock.Any(), classId).Return(class, true).AnyTimes()
	keeper.SetClassesByAccount(ctx, types.ClassesByAccount{
		Account:  ownerAddress,
		ClassIds: []string{classId},
	})
	nftKeeper.EXPECT().GetTotalSupply(ctx, classId).Return(uint64(totalSupply))
	nftKeeper.EXPECT().Mint(gomock.Any(), gomock.Any(), ownerAddressBytes).Return(nil).Times(mintToOwnerCount)
	nftKeeper.EXPECT().UpdateClass(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	for i := 0; i < mintableCount; i++ {
		keeper.SetMintableNFT(ctx, types.MintableNFT{
			ClassId: classId,
			Id:      fmt.Sprintf("mintable%d", i),
			Input: types.NFTInput{
				Uri: fmt.Sprintf("mintable%d", i),
			},
		})
	}
	nftKeeper.EXPECT().GetNFTsOfClass(ctx, classId).Return([]nft.NFT{})

	// call
	err = keeper.RevealMintableNFTs(ctx, classId)
	require.Error(t, err)
	require.Contains(t, err.Error(), "length mismatch")

	ctrl.Finish()
}

func TestRevealFailToUpdateToken(t *testing.T) {
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
	// hash from mainnet block 1
	hash, err := hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")
	require.NoError(t, err)
	ctx = ctx.WithBlockHeader(tmproto.Header{
		AppHash: hash,
	})

	ownerAddressBytes := []byte{0, 1, 0, 1, 0, 1, 0, 1}
	ownerAddress, _ := sdk.Bech32ifyAddressBytes("cosmos", ownerAddressBytes)
	classId := "likenft1aabbccddeeff"
	supply := 100
	mintableCount := 99
	totalSupply := 90
	mintToOwnerCount := mintableCount - totalSupply

	classData := types.ClassData{
		Metadata: types.JsonInput(`{"aaaa": "bbbb"}`),
		Parent: types.ClassParent{
			Type:    types.ClassParentType_ACCOUNT,
			Account: ownerAddress,
		},
		Config: types.ClassConfig{
			Burnable:  false,
			MaxSupply: uint64(supply),
			BlindBoxConfig: &types.BlindBoxConfig{
				MintPeriods: []types.MintPeriod{
					{
						StartTime:        time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
						AllowedAddresses: []string{},
						MintPrice:        uint64(0),
					},
				},
			},
		},
		MintableCount: uint64(mintableCount),
		ToBeRevealed:  true,
	}
	classDataInAny, _ := cdctypes.NewAnyWithValue(&classData)
	class := nft.Class{
		Id:          classId,
		Name:        "Class Name",
		Symbol:      "ABC",
		Description: "Testing Class 123",
		Uri:         "ipfs://abcdef",
		UriHash:     "abcdef",
		Data:        classDataInAny,
	}
	// mock calls
	nftKeeper.EXPECT().GetClass(gomock.Any(), classId).Return(class, true).AnyTimes()
	keeper.SetClassesByAccount(ctx, types.ClassesByAccount{
		Account:  ownerAddress,
		ClassIds: []string{classId},
	})
	nftKeeper.EXPECT().GetTotalSupply(ctx, classId).Return(uint64(totalSupply))
	nftKeeper.EXPECT().Mint(gomock.Any(), gomock.Any(), ownerAddressBytes).Return(nil).Times(mintToOwnerCount)
	nftKeeper.EXPECT().UpdateClass(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	for i := 0; i < mintableCount; i++ {
		keeper.SetMintableNFT(ctx, types.MintableNFT{
			ClassId: classId,
			Id:      fmt.Sprintf("mintable%d", i),
			Input: types.NFTInput{
				Uri: fmt.Sprintf("mintable%d", i),
			},
		})
	}
	var blindTokens []nft.NFT
	for i := 0; i < totalSupply+mintToOwnerCount; i++ {
		nftData := types.NFTData{
			ClassParent:  classData.Parent,
			ToBeRevealed: true,
		}
		nftDataInAny, err := cdctypes.NewAnyWithValue(&nftData)
		require.NoError(t, err)
		blindTokens = append(blindTokens, nft.NFT{
			ClassId: classId,
			Id:      fmt.Sprintf("nft%d", i),
			Data:    nftDataInAny,
		})
	}
	nftKeeper.EXPECT().GetNFTsOfClass(ctx, classId).Return(blindTokens)
	nftKeeper.EXPECT().Update(ctx, gomock.Any()).Return(fmt.Errorf("Failed to update"))

	// call
	err = keeper.RevealMintableNFTs(ctx, classId)
	require.Error(t, err)
	require.Contains(t, err.Error(), types.ErrFailedToUpdateNFT.Error())

	ctrl.Finish()
}
