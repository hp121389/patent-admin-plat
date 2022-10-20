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

type Patent struct {
	service.Service
}

// GetPage 获取patent列表
func (e *Patent) GetPage(c *dto.PatentGetPageReq, list *[]models.Patent, count *int64) error {
	var err error
	var data models.Patent

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

// Get 获取Patent对象
func (e *Patent) Get(d *dto.PatentById, model *models.Patent) error {
	//引用传递、函数名、形参、返回值
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

// Remove 根据专利id删除Patent
func (e *Patent) Remove(c *dto.PatentById) error {
	var err error
	var data models.Patent

	db := e.Orm.Delete(&data, c.GetPatentId())

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

// UpdateLists 根据PatentId修改Patent对象
func (e *Patent) UpdateLists(c *dto.PatentUpdateReq) error {
	var err error
	var model models.Patent
	db := e.Orm.First(&model, c.GetPatentId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update Patent error: %s", err)
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

// Insert 根据PatentId 创建Patent对象
func (e *Patent) Insert(c *dto.PatentInsertReq) error {
	var err error
	var data models.Patent
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

// InsertIfAbsent 根据PatentId 创建Patent对象
func (e *Patent) InsertIfAbsent(c *dto.PatentInsertReq) (int, error) {
	var err error
	var data models.Patent
	var i int64
	err = e.Orm.Model(&data).Where("PNM = ?", c.PNM).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}
	if i > 0 {
		err = e.Orm.Model(&data).Where("PNM = ?", c.PNM).First(&data).Error
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return 0, err
		}
		return data.PatentId, nil
	}
	c.GenerateList(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}
	return data.PatentId, nil
}
