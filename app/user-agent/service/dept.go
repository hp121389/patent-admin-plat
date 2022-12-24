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
