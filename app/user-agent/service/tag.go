package service

import (
	"errors"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
	"gorm.io/gorm"
)

type Tag struct {
	service.Service
}

// Get 获取SysRole对象
func (e *Tag) Get(d *dto.TagGetReq, model *models.Tag) error {
	var err error
	db := e.Orm.First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *Tag) Insert(d *dto.TagInsertReq) (int, error) {
	var err error
	var data models.Tag
	var i int64
	err = e.Orm.Model(&data).Where("tag_name = ?", d.TagName).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}
	if i > 0 {
		err := errors.New("用户名已存在！")
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}
	d.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}
	return data.TagId, nil
}

func (e *Tag) Remove(c *dto.TagById) error {
	var err error
	var data models.Tag

	db := e.Orm.Model(&data).
		Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveSysUser : %s", err)
		return err
	}
	//if db.RowsAffected == 0 {
	//	return errors.New("无权删除该数据")
	//}
	return nil
}

func (e *Tag) Update(c *dto.TagUpdateReq) error {
	var err error
	var model models.Tag
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	update := e.Orm.Model(&model).Where("tag_id = ?", &model.TagId).Omit("password", "salt").Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update userinfo error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

func (e Tag) NewUserTagInsert(c *dto.UserTagInsertReq) error {
	var err error
	var data models.UserTag
	var i int64
	err = e.Orm.Model(&data).Where("User_Id = ? and Tag_Id = ?", c.UserId, c.TagId).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
	}
	if i > 0 {
		err = fmt.Errorf("%w, (p:%d, u:%d) existed", ErrConflictBindPatent, c.UserId, c.TagId)
		e.Log.Errorf("db error: %s", err)
		return err
	}

	c.GenerateUserPatent(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}
