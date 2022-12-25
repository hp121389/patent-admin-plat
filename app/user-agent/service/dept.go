package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service/dtos"
)

type Dept struct {
	service.Service
}

// JoinDept 用户加入Dept申请
func (e *Dept) JoinDept(c *dtos.DeptReq) error {
	var err error
	var data model.Dept
	var i int64
	err = e.Orm.Model(&data).Where("dept_name = ?", c.DeptName).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("部门名称已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.GenerateDept(&data)
	data.UpdatedAt = dtos.UpdateTime()
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// GetDeptRelaListByUser 获取DeptRela对象列表通过deptId
func (e *Dept) GetDeptRelaListByUser(id int, list *[]model.DeptRelation) error {
	var err error
	var data []model.DeptRelation
	err = e.Orm.Model(&data).Where("user_id = ?", id).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetDeptUser 用户在该部门的状态
func (e *Dept) GetDeptUser(deptId int, userId int, list *[]model.DeptRelation) error {
	var err error
	var data []model.DeptRelation
	err = e.Orm.Model(&data).Where("dept_id = ? and user_id = ?", deptId, userId).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertDeptRela 添加DeptRela对象
func (e *Dept) InsertDeptRela(c *dtos.DeptRelaReq) error {
	var err error
	var data model.DeptRelation
	var i int64
	err = e.Orm.Model(&data).Where("dept_id = ? and user_id = ? ", c.DeptId, c.UserId).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("用户在该部门已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.GenerateRela(&data)
	data.UpdatedAt = dtos.UpdateTime()
	data.CreatedAt = dtos.UpdateTime()
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
