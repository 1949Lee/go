package server

import "leeBlogCli/definition"

// 获取登录的key(管理员email)
func (b *Blog) GetLoginKey() definition.GetLoginKeyResponse {
	author := b.Dao.GetAdminInfo()
	if author.IsActive == 0 {
		return definition.GetLoginKeyResponse{}
	}
	//res.Email = email.Email
	return definition.GetLoginKeyResponse{Email: author.Email}
}
