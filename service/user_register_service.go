package service

import (
	"VMP/model"
	"VMP/serializer"
)

// 管理用户注册的服务
type UserRegisterService struct {
	NickName 		string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName 		string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password 		string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Status: 	40001,
			Msg:	"两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("nickname = ?", service.NickName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 	40001,
			Msg: 	"昵称被占用",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 	40001,
			Msg: 	"用户已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		NickName:	service.NickName,
		UserName: 	service.UserName,
		Status: 	model.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return user, err
	}

	// 密码加密
	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Status:		40002,
			Msg:	"密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Status:		40002,
			Msg:		"密码加密失败",
		}
	}

	return user, nil
}