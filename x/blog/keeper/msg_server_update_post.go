package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"planet/x/blog/types"
)

func (k msgServer) SendUpdatePost(goCtx context.Context, msg *types.MsgSendUpdatePost) (*types.MsgSendUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.UpdatePostPacketData

	packet.PostID = msg.PostID
	packet.Title = msg.Title
	packet.Content = msg.Content

	// Transmit the packet
	err := k.TransmitUpdatePostPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendUpdatePostResponse{}, nil
}
