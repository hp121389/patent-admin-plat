package service

import (
	"errors"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
)

type SysDept struct {
	service.Service
}

// GetDeptList 获取Dept对象列表
func (e *SysDept) GetDeptList(list *[]models.SysDept) error {
	var err error
	var data []models.SysDept
	err = e.Orm.Model(&data).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	for _, d := range data {
		if err = e.CountMembers(&d); err != nil {
			return err
		}
	}
	return nil
}

// GetDeptByUserID 获取Dept by 用户id
func (e *SysDept) GetDeptByUserID(uid int) (*models.SysDept, error) {
	var err error
	var dr models.DeptRelation
	err = e.Orm.Model(&dr).
		Where("user_id = ?", uid).
		First(&dr).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}

	var dept models.SysDept
	err = e.Orm.Model(&dept).
		First(&dept, dr.DeptId).
		Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}

	if err = e.CountMembers(&dept); err != nil {
		return nil, err
	}

	return &dept, nil
}

func (e *SysDept) CountMembers(dept *models.SysDept) error {
	var err error
	var dr models.DeptRelation
	var count int64
	err = e.Orm.Model(&dr).
		Where("dept_id = ?", dept.DeptId).
		Count(&count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	dept.Members = int(count)
	return nil
}

// Insert 添加Dept对象
func (e *SysDept) Insert(c *dto.DeptReq) error {
	var err error
	var data models.SysDept
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
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *SysDept) Update(c *dto.DeptReq) error {
	var err error
	var model models.SysDept
	db := e.Orm.Where("Dept_Id = ?  ", c.DeptId).First(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update SysDept error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("dept not found")

	}
	c.GenerateDept(&model)

	update := e.Orm.Model(&model).Where("Dept_id = ?", &model.DeptId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update report-info error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// Remove 删除SysUser
func (e *SysDept) Remove(id int) error {
	var err error
	var data models.SysDept

	db := e.Orm.Model(&data).
		Delete(&data, id)
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveSysUser : %s", err)
		return err
	}
	return nil
}
