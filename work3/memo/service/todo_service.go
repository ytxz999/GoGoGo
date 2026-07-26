package service

import (
	"errors"
	"memo/dao"
	"memo/model"
	"time"
)

func CreateTodo(title string, content string, startTime string, endTime string, userId uint) error {

	//检验时间格式正确性
	//分钟必须是04
	StartTime, err := time.Parse("2006-01-02 15:04", startTime)
	if err != nil {
		return errors.New("开始时间格式错误")
	}

	EndTime, err := time.Parse("2006-01-02 15:04", endTime)
	if err != nil {
		return errors.New("结束时间格式错误")
	}

	//检验时间合理性
	if StartTime.After(EndTime) {
		return errors.New("开始时间不能大于结束时间")
	}

	//检查title是否为空
	if title == "" {
		return errors.New("标题不能为空")
	}

	//检查content是否为空
	if content == "" {
		return errors.New("内容不能为空")
	}

	//设置默认状态
	status := 0

	//检查时间是否正确

	//创建todo
	todo := model.Todo{
		Title:     title,
		Content:   content,
		StartTime: StartTime,
		EndTime:   EndTime,
		UserID:    userId,
		Status:    status,
	}

	return dao.CreateTodo(&todo)
}

func GetTodoList(page int, size int, status *int, keyword string, userId uint) ([]model.Todo, int64, error) {

	return dao.GetTodoList(page, size, status, keyword, userId)
}

func UpdateTodo(id *uint, status int, userId uint) error {
	if status != model.TodoStatusPending && status != model.TodoStatusDone {
		return errors.New("状态错误")
	}

	err := dao.UpdateTodo(id, status, userId)
	return err

}

func DeleteTodo(id *uint, status *int, all bool, userId uint) error {
	if status != nil {
		if *status != model.TodoStatusPending && *status != model.TodoStatusDone {
			return errors.New("状态错误")
		}
	}

	err := dao.DeleteTodo(id, status, all, userId)
	return err

}
