package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
)

type PatentPackage struct {
	service.Service
}

//GetPatentIdByPackageId 通过PackageId获得PatentId
func (e *PatentPackage) GetPatentIdByPackageId(c *dto.PackagePageGetReq, list *[]models.PatentPackage, count *int64) error {
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
func (e *PatentPackage) InsertPatentPackage(c *dto.PatentPackageReq) error {
	var err error
	var data models.PatentPackage
	var i int64
	err = e.Orm.Model(&data).Where("PNM = ? AND Package_Id = ? ", c.PNM, c.PackageId).
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

// RemovePackagePatent 根据专利和专利包id删除专利包关系
func (e *PatentPackage) RemovePackagePatent(c *dto.PackagePageGetReq) error {
	var err error
	var data models.PatentPackage

	db := e.Orm.Where("PNM = ? AND Package_Id = ? ", c.PNM, c.PackageId).
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

// IsPatentInPackage 判断专利是否在专利包中
func (e *PatentPackage) IsPatentInPackage(c *dto.PatentPackageReq) (bool, error) {
	var err error
	var data models.PatentPackage
	var i int64

	err = e.Orm.Model(&data).Where("PNM = ? AND Package_Id = ? ", c.PNM, c.PackageId).
		Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return false, err
	}
	if i > 0 {
		return true, nil
	}
	return false, nil
}
