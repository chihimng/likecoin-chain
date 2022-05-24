package e2e_test

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/likecoin/likechain/x/likenft/types"
)

func parseEventCreateClass(res sdk.TxResponse) types.EventNewClass {
	actualEvent := types.EventNewClass{}

ParseEventCreateClass:
	for _, log := range res.Logs {
		for _, event := range log.Events {
			if event.Type == "likechain.likenft.EventNewClass" {
				for _, attr := range event.Attributes {
					if attr.Key == "class_id" {
						actualEvent.ClassId = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "parent_iscn_id_prefix" {
						actualEvent.ParentIscnIdPrefix = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "parent_account" {
						actualEvent.ParentAccount = strings.Trim(attr.Value, "\"")
					}
				}
				break ParseEventCreateClass
			}
		}
	}

	return actualEvent
}

func parseEventUpdateClass(res sdk.TxResponse) types.EventUpdateClass {
	actualEvent := types.EventUpdateClass{}

ParseEventUpdateClass:
	for _, log := range res.Logs {
		for _, event := range log.Events {
			if event.Type == "likechain.likenft.EventUpdateClass" {
				for _, attr := range event.Attributes {
					if attr.Key == "class_id" {
						actualEvent.ClassId = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "parent_iscn_id_prefix" {
						actualEvent.ParentIscnIdPrefix = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "parent_account" {
						actualEvent.ParentAccount = strings.Trim(attr.Value, "\"")
					}
				}
				break ParseEventUpdateClass
			}
		}
	}

	return actualEvent
}

func parseEventMintNFT(res sdk.TxResponse) types.EventMintNFT {
	actualEvent := types.EventMintNFT{}

ParseEventMintNFT:
	for _, log := range res.Logs {
		for _, event := range log.Events {
			if event.Type == "likechain.likenft.EventMintNFT" {
				for _, attr := range event.Attributes {
					if attr.Key == "class_id" {
						actualEvent.ClassId = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "nft_id" {
						actualEvent.NftId = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "owner" {
						actualEvent.Owner = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "class_parent_iscn_id_prefix" {
						actualEvent.ClassParentIscnIdPrefix = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "class_parent_account" {
						actualEvent.ClassParentAccount = strings.Trim(attr.Value, "\"")
					}
				}
				break ParseEventMintNFT
			}
		}
	}

	return actualEvent
}

func parseEventBurnNFT(res sdk.TxResponse) types.EventBurnNFT {
	actualEvent := types.EventBurnNFT{}

ParseEventBurnNFT:
	for _, log := range res.Logs {
		for _, event := range log.Events {
			if event.Type == "likechain.likenft.EventBurnNFT" {
				for _, attr := range event.Attributes {
					if attr.Key == "class_id" {
						actualEvent.ClassId = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "nft_id" {
						actualEvent.NftId = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "owner" {
						actualEvent.Owner = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "class_parent_iscn_id_prefix" {
						actualEvent.ClassParentIscnIdPrefix = strings.Trim(attr.Value, "\"")
					}
					if attr.Key == "class_parent_account" {
						actualEvent.ClassParentAccount = strings.Trim(attr.Value, "\"")
					}
				}
				break ParseEventBurnNFT
			}
		}
	}

	return actualEvent
}