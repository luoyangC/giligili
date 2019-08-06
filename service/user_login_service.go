package service

import (
	"giligili/util"

	"giligili/model"
	"giligili/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// LoginResult 登录结果结构
type LoginResult struct {
	Token    string `json:"token"`
	UserName string
	NickName string
}

// Login 用户登录函数
func (service *UserLoginService) Login() (*serializer.Response) {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}
	
	res := generateToken(user)
	return res
}

// 生成令牌
func generateToken(user model.User) *serializer.Response {
    j := util.NewJWT()
    claims := util.NewCustomClaims(user.ID, user.UserName)

    token, err := j.CreateToken(claims)

    if err != nil {
		return &serializer.Response{
			Status: 4000,
			Msg: "登录失败",
		}
    }

    data := LoginResult{
		Token: token,
		UserName: user.UserName,
		NickName: user.Nickname,
	}
	
	return &serializer.Response{
		Status: 2000,
		Data: data,
	}
}