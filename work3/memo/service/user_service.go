package service

// 业务逻辑层
// 职责：负责核心业务规则计算、业务流程编排、事务边界控制及外部资源（RPC/缓存）协调。
import (
	"errors"
	"memo/dao"
	"memo/model"
	"memo/utils/jwt"

	"golang.org/x/crypto/bcrypt"
)

func Register(username string, password string) error {

	//用户注册,检查用户名是否为空
	if username == "" {
		return errors.New("用户名不能为空")
	}

	//检查用户是否存在
	if _, err := dao.FindUser(username); err == nil {
		return errors.New("用户已存在")
	}

	//用户密码加密
	//密码必须是一串字节切片，在底层处理加密和字符编码时更安全
	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return errors.New("用户密码加密失败")
	}

	//创建用户
	user := model.User{
		Username: username,
		Password: string(hashPassword),
	}

	return dao.CreateUser(&user)
}

func Login(username string, password string) (string, error) {
	//查询用户
	user, err := dao.FindUser(username)
	if err != nil {
		return "", errors.New("用户不存在")
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return "", errors.New("密码错误")
	}

	//生成token
	return utils.GenerateToken(user.Id)
}
