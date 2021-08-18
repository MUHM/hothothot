package logic

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"hothothot/common/errorx"
	"hothothot/common/tools/captcha"
	baseTypes "hothothot/common/types"
	"hothothot/service/gateway/internal/svc"
	"hothothot/service/gateway/internal/types"
	"hothothot/service/system/rpc/systemclient"
	"hothothot/service/user/rpc/userclient"

	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {
	captchaRes := captcha.VerifyString(req.CaptchaId, req.Captcha)
	if !captchaRes {
		return nil, errorx.NewDefaultError("验证码错误")
	}
	userInfo, err := l.svcCtx.UserRpc.GetUserByName(l.ctx, &userclient.NameReq{Name: req.Username})
	passwordSetting, _ := l.svcCtx.SystemRpc.GetByName(l.ctx, &systemclient.NameReq{Name: "password"})
	var passwordConfig baseTypes.PasswordConfig
	json.Unmarshal([]byte(passwordSetting.Content), &passwordConfig)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	password := fmt.Sprintf("%x", md5.Sum([]byte(req.Password+passwordConfig.Secret)))
	if userInfo.Password != password {
		return nil, errorx.NewDefaultError("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id, userInfo.Username)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Id:           userInfo.Id,
		Username:     userInfo.Username,
		Nickname:     userInfo.Nickname,
		Gender:       userInfo.Gender,
		Mail:         userInfo.Mail,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64, name string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	claims["name"] = name
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
