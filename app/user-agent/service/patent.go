package service

import (
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"

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

// GetUserPatentPage 获取patent列表
func (e *Patent) GetUserPatentPage(c *dto.UserPatentObject, list *[]models.Patent, count *int64) error {
	var err error
	var data models.Patent

	err = e.Orm.Model(&data).Where("user_id = ?", c.UserId).
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
	db := e.Orm.First(&model, c.PatentId)
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
	err = e.Orm.Model(&data).Where("PNM = ?", c.PNM).Count(&i).Error
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

// InsertIfAbsent 根据PatentId 创建Patent对象 且返回创建对象的PatentId
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

// RemoveClaim 取消认领
func (e *Patent) RemoveClaim(c *dto.UserPatentObject) error {
	var err error
	var data models.UserPatent

	db := e.Orm.Where("Patent_Id = ? AND User_Id = ? AND Type = ?", c.PatentId, c.UserId, dto.ClaimType).
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

// RemoveFocus 取消关注
func (e *Patent) RemoveFocus(c *dto.UserPatentObject) error {
	var err error
	var data models.UserPatent

	db := e.Orm.Where("Patent_Id = ? AND User_Id = ? AND Type = ?", c.PatentId, c.UserId, dto.FocusType).
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

//// UpdateUserPatent
//func (e *UserPatent) UpdateUserPatent(c *dto.UpDateUserPatentObject) error {
//	var err error
//	var model models.UserPatent
//	var i int64
//
//	ids := e.Orm.Model(&model).Where("Patent_Id = ? AND User_Id = ? ", c.PatentId, c.UserId).First(&model).Count(&i)
//
//	fmt.Println("一共有", i, "个专利id为", c.PatentId, "且用户是", c.UserId, "的关系")
//
//	if i == 2 {
//		//先按照条件找到用户对应的专利，然后修改，且只找一个。
//		//如果一个用户即关注又认领了一个专利怎么办呢 ,model不是数组，只是一个model
//		return errors.New("您已同时认领和关注该专利！")
//	}
//
//	err = ids.Error
//
//	db := e.Orm.Model(&model).Where("Patent_Id = ? AND User_Id = ? ", c.PatentId, c.UserId).
//		First(&model)
//
//	if err = db.Error; err != nil {
//		e.Log.Errorf("Service Update User-Patent error: %s", err)
//		return err
//	}
//	if db.RowsAffected == 0 {
//		return errors.New("无权更新该数据")
//	}
//
//	c.GenerateUserPatent(&model)
//
//	update := e.Orm.Model(&model).Updates(&model)
//	if err = update.Error; err != nil {
//		e.Log.Errorf("db error: %s", err)
//		return err
//	}
//	if update.RowsAffected == 0 {
//		err = errors.New("update patent-info error maybe you dont need update or record not exist")
//		log.Warnf("db update error")
//		return err
//	}
//	return nil
//}

// GetClaimLists 通过UserId获得PatentId列表
func (e *Patent) GetClaimLists(c *dto.UserPatentGetPageReq, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("Type = ? AND User_Id = ?", dto.ClaimType, c.GetUserId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetFocusLists 通过UserId获得PatentId列表
func (e *Patent) GetFocusLists(c *dto.UserPatentGetPageReq, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("Type = ? AND User_Id = ?", dto.FocusType, c.GetUserId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetUserPatentIds 通过UserId获得PatentId列表
func (e *Patent) GetUserPatentIds(c *dto.UserPatentGetPageReq, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("User_Id = ?", c.GetUserId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetPatentPagesByIds 获取patent列表
func (e *Patent) GetPatentPagesByIds(d *dto.PatentsIds, list *[]models.Patent, count *int64) error {
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

// Inserts relationship between user and patent
func (e *Patent) Inserts(c *dto.UserPatentObject) error {
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

// GetTagIdByPatentId 通过PatentId获得TagId
func (e *Patent) GetTagIdByPatentId(c *dto.PatentTagGetPageReq, list *[]models.PatentTag, count *int64) error {
	var err error
	var data models.PatentTag

	err = e.Orm.Model(&data).
		Where("Patent_Id = ?", c.GetPatentId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetTagPages 通过TagId获取Tag列表（TagName等）
func (e *Patent) GetTagPages(d *dto.TagsByIdsForRelationshipPatents, list *[]models.Tag, count *int64) error {

	var err error
	var ids []int = d.GetTagId()

	for i := 0; i < len(ids); i++ {

		if ids[i] != 0 {

			var data1 models.Tag

			err = e.Orm.Model(&data1).
				Where("Tag_Id = ? ", ids[i]).
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

//GetPatentIdByTagId 通过TagId获得PatentId
func (e *Patent) GetPatentIdByTagId(c *dto.TagPageGetReq, list *[]models.PatentTag, count *int64) error {
	var err error
	var data models.PatentTag

	err = e.Orm.Model(&data).
		Where("Tag_Id = ?", c.GetTagId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetPatentPages 通过PatentId获取Patent列表（TI等）
func (e *Patent) GetPatentPages(d *dto.PatentsIds, list *[]models.Patent, count *int64) error {

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

// InsertPatentTagRelationship 创建专利标签关系
func (e *Patent) InsertPatentTagRelationship(c *dto.PatentTagInsertReq) error {
	var err error
	var data models.PatentTag
	var i int64
	err = e.Orm.Model(&data).Where("Patent_Id = ? AND Tag_Id = ? ", c.PatentId, c.TagId).
		Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("关系已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}

	c.GeneratePatentTag(&data)

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// RemoveRelationship 根据专利id、TYPE删除用户专利关系
func (e *Patent) RemoveRelationship(c *dto.PatentTagInsertReq) error {
	var err error
	var data models.PatentTag

	db := e.Orm.Where("Patent_Id = ? AND Tag_Id = ? ", c.PatentId, c.TagId).
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

//GetPatentIdByPackageId 通过PackageId获得PatentId
func (e *Patent) GetPatentIdByPackageId(c *dto.PackagePageGetReq, list *[]models.PatentPackage, count *int64) error {
	var err error
	var data models.PatentPackage

	err = e.Orm.Model(&data).
		Where("Package_Id = ?", c.PackageId).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertPatentPackage 创建专利标签关系
func (e *Patent) InsertPatentPackage(c *dto.PackagePageGetReq) error {
	var err error
	var data models.PatentPackage
	var i int64
	err = e.Orm.Model(&data).Where("Patent_Id = ? AND Package_Id = ? ", c.PatentId, c.PackageId).
		Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("关系已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}

	c.GeneratePackagePatent(&data)

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// RemovePackagePatent 根据专利id、TYPE删除用户专利关系
func (e *Patent) RemovePackagePatent(c *dto.PackagePageGetReq) error {
	var err error
	var data models.PatentPackage

	db := e.Orm.Where("Patent_Id = ? AND Package_Id = ? ", c.PatentId, c.PackageId).
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