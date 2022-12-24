package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"scavenge/x/scavenge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) RevealSolution(goCtx context.Context, msg *types.MsgRevealSolution) (*types.MsgRevealSolutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var solutionScavengerBytes = []byte(msg.Solution + msg.Creator)

	var solutionScavengerHash = sha256.Sum256(solutionScavengerBytes)

	var solutionScavengerHashString = hex.EncodeToString(solutionScavengerHash[:])
	var scavenge types.Scavenge

	_, isFound := k.GetCommit(ctx, solutionScavengerHashString)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Scavenge with that solution hash doesn't exists")
	}

	_, err := sdk.AccAddressFromBech32(scavenge.Scavenger)
	if err == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Scavenge has already been solved")
	}

	scavenge.Scavenger = msg.Creator

	scavenge.Solution = msg.Solution

	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	scavenger, err := sdk.AccAddressFromBech32(scavenge.Scavenger)
	if err != nil {
		panic(err)
	}

	reward, err := sdk.ParseCoinsNormalized(scavenge.Reward)
	if err != nil {
		panic(err)
	}

	sdkError := k.bankKeeper.SendCoins(ctx, moduleAcct, scavenger, reward)
	if sdkError != nil {
		return nil, sdkError
	}

	k.SetScavenge(ctx, scavenge)
	return &types.MsgRevealSolutionResponse{}, nil
}
