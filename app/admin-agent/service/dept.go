package service

import (
	"errors"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service/dtos"
)

type Dept struct {
	service.Service
}

// GetDeptList 获取Dept对象列表
func (e *Dept) GetDeptList(list *[]model.Dept) error {
	var err error
	var data []model.Dept
	err = e.Orm.Model(&data).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetDeptRelaList 获取Dept对象列表
func (e *Dept) GetDeptRelaList(list *[]model.DeptRelation) error {
	var err error
	var data []model.DeptRelation
	err = e.Orm.Model(&data).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetDeptRelaListById 获取Dept对象列表
func (e *Dept) GetDeptRelaListById(id int, list *[]model.DeptRelation) error {
	var err error
	var data []model.DeptRelation
	err = e.Orm.Model(&data).Where("dept_id = ?", id).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertDept 添加Dept对象
func (e *Dept) InsertDept(c *dtos.DeptReq) error {
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

// UpdateDept 更新Dept对象
func (e *Dept) UpdateDept(c *dtos.DeptReq) error {
	var err error
	var model model.Dept
	db := e.Orm.Where("Dept_Id = ?  ", c.DeptId).First(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update Dept error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("dept not found")

	}
	c.GenerateDept(&model)
	model.UpdatedAt = dtos.UpdateTime()

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

// UpdateDeptRela 更新DeptRela对象
func (e *Dept) UpdateDeptRela(c *dtos.DeptRelaReq) error {
	var err error
	var model model.DeptRelation
	db := e.Orm.Where("Dept_Id = ?  ", c.DeptId).First(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update Dept error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("dept not found")

	}
	c.GenerateRela(&model)
	model.UpdatedAt = dtos.UpdateTime()

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

// UserDeptRela 更新DeptRela对象
func (e *Dept) UserDeptRela(c *dtos.DeptRelaReq) error {
	var err error
	var model model.DeptRelation
	db := e.Orm.Where("Dept_Id = ? and User_id = ? ", c.DeptId, c.UserId).First(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update Dept error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("dept not found")

	}
	c.GenerateRela(&model)
	model.UpdatedAt = dtos.UpdateTime()

	update := e.Orm.Model(&model).Where("Dept_id = ? and User_id = ?", &model.DeptId, &model.UserId).Updates(&model)
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
