package service

import (
	"errors"
	"github.com/prometheus/common/log"
	"go-admin/app/patent/models"
	"go-admin/app/patent/service/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	cDto "go-admin/common/dto"
)

type SysList struct {
	service.Service
}

// GetPage 获取SysList列表
// service调用dto和models
func (e *SysList) GetPage(c *dto.SysListGetPageReq, list *[]models.SysList, count *int64) error {
	var err error
	var data models.SysList

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Get 获取SysList对象
func (e *SysList) Get(d *dto.SysListById, model *models.SysList) error {
	var err error
	db := e.Orm.First(model, d.GetPatentId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看专利不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

//// Get 获取SysList对象
//func (e *SysList) GetByName(d *dto.SysListByName, model *models.SysList) error {
//	var err error
//	db := e.Orm.First(model, d.GetPatentTI())
//	err = db.Error
//	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
//		err = errors.New("查看专利不存在或无权查看")
//		e.Log.Errorf("db error:%s", err)
//		return err
//	}
//	if db.Error != nil {
//		e.Log.Errorf("db error:%s", err)
//		return err
//	}
//	return nil
//}

// Remove 删除SysList
func (e *SysList) Remove(c *dto.SysListDeleteReq) error {
	var err error
	var data models.SysList

	db := e.Orm.Delete(&data, c.GetPatentId()).Where("patent_id = ?", c.GetPatentId())
	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// UpdateLists 根据PatentId修改SysList对象
func (e *SysList) UpdateLists(c *dto.SysListUpdateReq) error {
	var err error
	var model models.SysList
	db := e.Orm.First(&model, c.GetPatentId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysList error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.GenerateList(&model)
	update := e.Orm.Model(&model).Where("patent_id = ?", &model.PatentId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update patent-info error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// InsertListsByPatentId 根据PatentId 创建SysList对象
func (e *SysList) InsertListsByPatentId(c *dto.SysListInsertReq) error {
	var err error
	var data models.SysList
	var i int64
	err = e.Orm.Model(&data).Where("patent_id = ?", c.PatentId).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("专利ID已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.GenerateList(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
