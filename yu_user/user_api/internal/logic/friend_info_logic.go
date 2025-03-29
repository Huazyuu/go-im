package logic

import (
	"context"
	"encoding/json"
	"errors"
	"server/yu_user/user_models"
	"server/yu_user/user_rpc/types/user_rpc"

	"server/yu_user/user_api/internal/svc"
	"server/yu_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendInfoLogic {
	return &FriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendInfoLogic) FriendInfo(req *types.FriendInfoRequest) (resp *types.FriendInfoResponse, err error) {
	var fr user_models.FriendsModel
	if !fr.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("他不是你的好友哦~")
	}
	res, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{UserId: int32(req.FriendID)})
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var friend user_models.UserModel
	json.Unmarshal(res.Data, &friend)

	resp = &types.FriendInfoResponse{
		UserID:   friend.ID,
		Username: friend.UserName,
		Nickname: friend.UserNickname,
		Abstract: friend.UserAbstract,
		Avatar:   friend.UserAvatar,
		Notice:   fr.GetUserNotice(friend.ID),
	}
	return resp, nil
}
