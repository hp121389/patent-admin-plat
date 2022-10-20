package service

import (
	"errors"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/patent/models"
	"go-admin/app/patent/service/dto"
)

type UserPatent struct {
	service.Service
}

// GetClaimLists 通过UserId获得PatentId列表
func (e *UserPatent) GetClaimLists(c *dto.UserPatentGetPageReq, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("Type = ? AND User_Id = ?", "认领", c.GetUserId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetCollectionLists 通过UserId获得PatentId列表
func (e *UserPatent) GetCollectionLists(c *dto.UserPatentGetPageReq, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("Type = ? AND User_Id = ?", "关注", c.GetUserId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetPatentPagesByIds 获取patent列表
func (e *UserPatent) GetPatentPagesByIds(d *dto.PatentsByIdsForRelationshipUsers, list *[]models.Patent, count *int64) error {
	var err error
	var ids []int = d.GetPatentId()
	for i := 0; i < len(ids); i++ {
		if ids[i] != 0 {
			var data1 models.Patent
			err = e.Orm.Model(&data1).
				Where("Patent_Id = ? ", ids[i]).
				First(&data1).Limit(-1).Offset(-1).
				Count(count).Error
			*list = append(*list, data1)
			if err != nil {
				e.Log.Errorf("db error:%s", err)
				return err
			}
		}
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert relationship between user and patent
func (e *UserPatent) Insert(c *dto.UserPatentObject) error {
	var err error
	var data models.UserPatent
	var i int64
	err = e.Orm.Model(&data).Where("Patent_Id = ? AND User_Id = ? AND Type = ?", c.PatentId, c.UserId, c.Type).
		Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err = fmt.Errorf("%w, (p:%d, u:%d, t:%s) existed", ErrConflictBindPatent, c.PatentId, c.UserId, c.Type)
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

//// InsertCollectionRelationship 创建关注关系
//func (e *UserPatent) InsertCollectionRelationship(c *dto.UserPatentObject) error {
//	var err error
//	var data models.UserPatent
//	var i int64
//	err = e.Orm.Model(&data).Where("Patent_Id = ? AND User_Id = ? AND Type = ?", c.PatentId, c.UserId, c.Type).
//		Count(&i).Error
//	if err != nil {
//		e.Log.Errorf("db error: %s", err)
//		return err
//	}
//	if i > 0 {
//		err := errors.New("关系已存在！")
//		e.Log.Errorf("db error: %s", err)
//		return err
//	}
//
//	c.GenerateUserPatent(&data)
//	c.Type = "关注"
//
//	err = e.Orm.Create(&data).Error
//	if err != nil {
//		e.Log.Errorf("db error: %s", err)
//		return err
//	}
//	return nil
//}

// RemoveRelationship 根据专利id、TYPE删除用户专利关系
func (e *UserPatent) RemoveRelationship(c *dto.UserPatentObject) error {
	var err error
	var data models.UserPatent

	db := e.Orm.Where("Patent_Id = ? AND User_Id = ? AND Type = ?", c.PatentId, c.UserId, c.Type).
		Delete(&data)

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

// UpdateUserPatent 根据PatentId修改Patent对象
func (e *UserPatent) UpdateUserPatent(c *dto.UpDateUserPatentObject) error {
	var err error
	var model models.UserPatent
	var i int64

	ids := e.Orm.Model(&model).Where("Patent_Id = ? AND User_Id = ? ", c.PatentId, c.UserId).First(&model).Count(&i)

	fmt.Println("一共有", i, "个专利id为", c.PatentId, "且用户是", c.UserId, "的关系")

	if i == 2 {
		//先按照条件找到用户对应的专利，然后修改，且只找一个。
		//如果一个用户即关注又认领了一个专利怎么办呢 ,model不是数组，只是一个model
		return errors.New("您已同时认领和关注该专利！")
	}

	err = ids.Error

	db := e.Orm.Model(&model).Where("Patent_Id = ? AND User_Id = ? ", c.PatentId, c.UserId).
		First(&model)

	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update User-Patent error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	c.GenerateUserPatent(&model)

	update := e.Orm.Model(&model).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update patent-info error maybe you dont need update or record not exist")
		log.Warnf("db update error")
		return err
	}
	return nil
}
