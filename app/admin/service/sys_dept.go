package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
)

type SysDept struct {
	service.Service
}

// Insert a SysDept object
func (e *SysDept) Insert(c *dto.SysDeptInsertReq) error {
	var err error
	var data models.SysDept
	var i int64
	//c.Generate(&data)

	err = e.Orm.Model(&data).Where("dept_name = ?", c.DeptName).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error : %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("该部门名称已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error :%s", err)
		return err
	}
	return nil
}
