package dao

import (
	"errors"
	"memo/database"
	"memo/model"
)

// 创建todo
func CreateTodo(todo *model.Todo) error {
	return database.DB.Create(todo).Error
}

// 查询todo
func GetTodoList(page *int, size int, status *int, keyword string, user_id uint) ([]model.Todo, int64, error) {
	//统计数量
	var total int64
	//限制用户
	query := database.DB.Model(&model.Todo{}).Where("user_id = ?", user_id)
	var truePage int
	if page == nil {
		truePage = 1
	} else {
		truePage = *page
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if keyword != "" {
		query = query.Where("title like ? or content like ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	var todos []model.Todo
	query.Count(&total)
	//实现分页查询
	err := query.Limit(size).Offset((truePage - 1) * size).Find(&todos).Error
	return todos, total, err
}

// 修改todo
func UpdateTodo(id *uint, status int, userId uint) error {
	//限制用户
	query := database.DB.Model(&model.Todo{}).Where("user_id = ?", userId)

	if id != nil {
		query = query.Where("id = ?", *id)
	}
	result := query.Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("没有修改任何数据")
	}
	return nil
}

// 删除todo
func DeleteTodo(id *uint, status *int, all bool, userId uint) error {
	//限制用户
	query := database.DB.Model(&model.Todo{}).Where("user_id = ?", userId)
	//可以将if改成switch，service传入判断删除方式的值
	//后面再改
	//dao层判断有点多，感觉在service层
	if id != nil && status == nil && !all {
		//删除单个备忘录
		query = query.Where("id = ?", *id)

	} else if id == nil && status != nil && !all {
		//删除某状态的事项
		query = query.Where("status = ?", *status)

	} else if id == nil && status == nil && all {
		//删除所有事项

	} else {
		return errors.New("参数错误")
	}
	result := query.Delete(&model.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("没有删除任何数据")
	}
	return nil
}
