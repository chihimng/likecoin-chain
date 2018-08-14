package handlers

import (
	"testing"

	"github.com/likecoin/likechain/abci/context"
	"github.com/likecoin/likechain/abci/types"

	. "github.com/smartystreets/goconvey/convey"
)

func wrapRegisterTransaction(tx *types.RegisterTransaction) *types.Transaction {
	return &types.Transaction{
		Tx: &types.Transaction_RegisterTx{
			RegisterTx: tx,
		},
	}
}

func TestCheckAndDeliverRegister(t *testing.T) {
	ctx := context.NewMock()
	Convey("Given a Register Transaction", t, func() {
		Convey("If it is a valid Register transaction", func() {
			ctx.Reset()
			rawTx := wrapRegisterTransaction(&types.RegisterTransaction{
				// TODO
			})
			Convey("CheckTx should return Code 0", func() {
				res := checkRegister(ctx, rawTx)
				So(res.Code, ShouldEqual, 0)
			})
			Convey("DeliverTx should return Code 0", func() {
				res := deliverRegister(ctx, rawTx)
				So(res.Code, ShouldEqual, 0)
				Convey("Should be able to query user account info afterwards", func() {
					_ = res.Data // TODO: ID
					// TODO: query
				})
			})
		})
		Convey("If it is a Register transaction with invalid address format", func() {
			ctx.Reset()
			rawTx := wrapRegisterTransaction(&types.RegisterTransaction{
				// TODO
			})
			Convey("CheckTx should return Code 1001", func() {
				res := checkRegister(ctx, rawTx)
				So(res.Code, ShouldEqual, 1001)
			})
			Convey("DeliverTx should return Code 1001", func() {
				res := deliverRegister(ctx, rawTx)
				So(res.Code, ShouldEqual, 1001)
			})
		})
		Convey("If it is a Register transaction with invalid signature version", func() {
			ctx.Reset()
			rawTx := wrapRegisterTransaction(&types.RegisterTransaction{
				// TODO
			})
			Convey("CheckTx should return Code 1001", func() {
				res := checkRegister(ctx, rawTx)
				So(res.Code, ShouldEqual, 1001)
			})
			Convey("DeliverTx should return Code 1001", func() {
				res := deliverRegister(ctx, rawTx)
				So(res.Code, ShouldEqual, 1001)
			})
		})
		Convey("If it is a Register transaction with invalid signature format", func() {
			ctx.Reset()
			rawTx := wrapRegisterTransaction(&types.RegisterTransaction{
				// TODO
			})
			Convey("CheckTx should return Code 1001", func() {
				res := checkRegister(ctx, rawTx)
				So(res.Code, ShouldEqual, 1001)
			})
			Convey("DeliverTx should return Code 1001", func() {
				res := deliverRegister(ctx, rawTx)
				So(res.Code, ShouldEqual, 1001)
			})
		})
		Convey("If it is a Register transaction with address already registered", func() {
			ctx.Reset()
			rawTx := wrapRegisterTransaction(&types.RegisterTransaction{
				// TODO
			})
			deliverRegister(ctx, rawTx)
			Convey("CheckTx should return Code 1002", func() {
				res := checkRegister(ctx, rawTx)
				So(res.Code, ShouldEqual, 1002)
			})
			Convey("DeliverTx should return Code 1002", func() {
				res := deliverRegister(ctx, rawTx)
				So(res.Code, ShouldEqual, 1002)
			})
		})
	})
}

func TestValidateRegisterSignature(t *testing.T) {
	Convey("Given a valid Register transaction", t, func() {
		ctx := context.NewMock()
		tx := &types.RegisterTransaction{} // TODO
		Convey("It should be able to register", func() {
			So(validateRegisterSignature(ctx, tx), ShouldBeTrue)
		})
		Convey("But should not be able to register again", func() {
			So(validateRegisterSignature(ctx, tx), ShouldBeFalse)
		})
	})
}

func TestValidateRegisterTransaction(t *testing.T) {
	Convey("Given a Register transaction", t, func() {
		Convey("With valid format", func() {
			tx := &types.RegisterTransaction{} // TODO
			Convey("It should pass the validation", func() {
				So(validateRegisterTransaction(tx), ShouldBeTrue)
			})
		})
		Convey("With invalid address format", func() {
			tx := &types.RegisterTransaction{} // TODO
			Convey("It should fail the validation", func() {
				So(validateRegisterTransaction(tx), ShouldBeFalse)
			})
		})
		Convey("With invalid signature version", func() {
			tx := &types.RegisterTransaction{} // TODO
			Convey("It should fail the validation", func() {
				So(validateRegisterTransaction(tx), ShouldBeFalse)
			})
		})
		Convey("With invalid signature format", func() {
			tx := &types.RegisterTransaction{} // TODO
			Convey("It should fail the validation", func() {
				So(validateRegisterTransaction(tx), ShouldBeFalse)
			})
		})
	})
}

func TestRegister(t *testing.T) {
	// TODO
}