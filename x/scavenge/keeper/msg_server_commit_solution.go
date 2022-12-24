package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CommitSolution(goCtx context.Context, msg *types.MsgCommitSolution) (*types.MsgCommitSolutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var commit = types.Commit{
		Index:                 msg.SolutionScavengerHash,
		SolutionHash:          msg.SolutionHash,
		SolutionScavengerHash: msg.SolutionScavengerHash,
	}

	_, isFound := k.GetCommit(ctx, commit.SolutionScavengerHash)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Commit with that hash already exists")
	}

	k.SetCommit(ctx, commit)

	return &types.MsgCommitSolutionResponse{}, nil
}
