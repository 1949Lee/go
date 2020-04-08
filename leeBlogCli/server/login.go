package server

import (
	"crypto/md5"
	"fmt"
	"leeBlogCli/definition"
	"strings"
)

// 获取登录的key(管理员email)
func (b *Blog) GetLoginKey() definition.GetLoginKeyResponse {
	author := b.Dao.GetAdminInfo()
	if author.IsActive == 0 {
		return definition.GetLoginKeyResponse{}
	}
	//res.Email = email.Email
	return definition.GetLoginKeyResponse{Email: author.Email}
}

// 根据传入的登录（扫描二维码）信息进行认证
func (b *Blog) ConfirmLoginInfo(param *definition.ConfirmLoginParam) definition.ConfirmLoginResponse {
	user := b.Dao.GetAuthorByEmail(param.Email)

	if user.ID < 0 {
		return definition.ConfirmLoginResponse{LeeToken: "-1"}
	}

	// email和设备的uuid拼接后的md5串
	leeToken := md5.Sum([]byte(user.Email + user.DeviceUUID))

	// 将得到的leeToken取出，update进数据库
	tokenStr := fmt.Sprintf("%X", leeToken)
	if tokenStr == strings.ToUpper(param.Passport) {
		err := b.Dao.UpdateAuthorToken(user.Email, tokenStr)
		if err == nil {
			return definition.ConfirmLoginResponse{LeeToken: tokenStr}
		}
	}
	return definition.ConfirmLoginResponse{LeeToken: "-1"}
}
