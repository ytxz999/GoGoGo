package dao

// 数据访问层
// 职责：负责数据库的 CRUD 原子操作、ORM 映射、基础查询条件构建。
import (
	"memo/database"
	"memo/model"
)

// 注册用户
func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

// 查找用户
func FindUser(username string) (*model.User, error) {
	var user model.User
	//不要用字符串拼接的sql语句，容易受到sql注入
	err := database.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
