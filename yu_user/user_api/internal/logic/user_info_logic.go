package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"server/yu_user/user_models"
	"server/yu_user/user_rpc/types/user_rpc"

	"server/yu_user/user_api/internal/svc"
	"server/yu_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	rpcRes, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{UserId: int32(req.UserID)})
	if err != nil {
		logx.Error("系统错误")
		return nil, err
	}
	var user user_models.UserModel
	err = json.Unmarshal(rpcRes.Data, &user)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据错误")
	}
	fmt.Println(user)
	resp = &types.UserInfoResponse{
		UserID:               user.ID,
		UserName:             user.UserName,
		UserNickname:         user.UserNickname,
		UserAbstract:         user.UserAbstract,
		UserAvatar:           user.UserAvatar,
		RecallMessage:        user.UserConfModel.RecallMsg,
		IsFriendOnlineNotify: user.UserConfModel.IsFriendOnlineNotify,
		IsSound:              user.UserConfModel.IsSound,
		IsSecureLink:         user.UserConfModel.IsSecureLink,
		IsSavePwd:            user.UserConfModel.IsSavePwd,
		SearchUser:           user.UserConfModel.SearchUser,
		Verification:         user.UserConfModel.Verification,
	}
	if user.UserConfModel.VerificationQuestion != nil {
		resp.VerificationQuestion = &types.VerificationQuestion{
			Question1: user.UserConfModel.VerificationQuestion.Question1,
			Question2: user.UserConfModel.VerificationQuestion.Question2,
			Question3: user.UserConfModel.VerificationQuestion.Question3,
			Answer1:   user.UserConfModel.VerificationQuestion.Answer1,
			Answer2:   user.UserConfModel.VerificationQuestion.Answer2,
			Answer3:   user.UserConfModel.VerificationQuestion.Answer3,
		}
	}
	return
}
