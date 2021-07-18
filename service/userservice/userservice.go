package userservice

import (
	"boframe/dao/userdao"
	"boframe/models/usermodels"
	"boframe/pkg/errcode"
	"boframe/pkg/jwt"
	"boframe/pkg/snowflake"
	"boframe/settings/mysqlI"
	"go.uber.org/zap"
)

type RegisterBo struct {
	Username   string
	Nickname   string
	Email      string
	Password   string
	RePassword string
}

// 传不传Context好呢?
func Register(bo *RegisterBo) (uid int64, err error) {
	// 判断用户名是否已经存在
	exists, err := userdao.CheckUsernameExists(bo.Username)
	if err != nil && !mysqlI.IsErrNotFound(err) {
		return 0, errcode.Build(errcode.ErrDB)
	}
	if exists {
		return 0, errcode.Build(errcode.ErrRegisterUsernameExists)
	}
	
	// 生成业务ID
	uid = snowflake.GenId()
	
	var hash string
	// 密码的加密问题
	if hash, err = usermodels.Encrypt(bo.Password); err != nil {
		return 0, errcode.Build(errcode.ErrUnknown)
	}
	
	// 插入数据库
	user := usermodels.User{
		UserId:   uid,
		Username: bo.Username,
		NickName: bo.Nickname,
		Password: hash,
		Email:    bo.Email,
	}
	if err = userdao.Insert(&user); err != nil {
		return 0, errcode.Build(errcode.ErrDB)
	}
	return uid, nil
}

type LoginBo struct {
	Username string
	Password string
}

func Login(bo *LoginBo) (token string, err error) {
	// 根据账号查找出信息
	user, err := userdao.OneByUsernameNotDeleted(bo.Username)
	if mysqlI.IsErrNotFound(err) {
		return "", errcode.Build(errcode.ErrLoginWrongPassword)
	}
	if err != nil {
		zap.L().Error("login service failed",
			zap.String("username", bo.Username),
			zap.Error(err),
		)
		return "", errcode.Build(errcode.ErrDB)
	}
	
	// 做密码的等比
	if ok := usermodels.Decrypt(user.Password, bo.Password); !ok {
		return "", errcode.Build(errcode.ErrLoginWrongPassword)
	}
	
	// 登录成功后生成token返回
	token, err = jwt.GenToken(user.UserId, user.Username)
	if err != nil {
		return "", errcode.Build(errcode.ErrServer)
	}
	return
}
