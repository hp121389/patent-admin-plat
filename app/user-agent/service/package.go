package service

import (
	"errors"
	"fmt"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	cDto "go-admin/common/dto"
)

type Package struct {
	service.Service
}

// GetPage 获取Package列表
func (e *Package) GetPage(c *dto.PackageGetPageReq, list *[]models.Package, count *int64) error {
	var err error
	//var data models.Package
	// todo: check
	err = e.Orm.Debug().
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			//actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Get 获取Package对象
func (e *Package) Get(d *dto.PackageById, model *models.Package) error {
	var data models.Package

	err := e.Orm.Model(&data).Debug().
		//Scopes(
		//	actions.Permission(data.TableName(), p),
		//).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Insert 创建Package对象
func (e *Package) Insert(c *dto.PackageInsertReq) (int, error) {
	var err error
	var data models.Package
	c.GenerateList(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return 0, err
	}
	return data.PackageId, nil
}

// Update 修改Package对象
func (e *Package) Update(c *dto.PackageUpdateReq) error {
	var err error
	var model models.Package
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	update := e.Orm.Model(&model).Where("package_id = ?", &model.PackageId).Updates(&model)
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

// Remove 删除Package
func (e *Package) Remove(c *dto.PackageById) error {
	var err error
	var data models.Package

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

type UserPackage struct {
	service.Service
}

// GetPage 获取user-package列表
func (e *UserPackage) GetPage(c *dto.UserPackageGetPageReq, list *[]models.UserPackage, count *int64) error {
	var err error
	var data models.UserPackage

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

// GetPackageIdsByUserId 获取User对应PackageId列表
func (e *UserPackage) GetPackageIdsByUserId(d *dto.UserPackageGetPageReq, list *[]models.UserPackage, count *int64) error {
	var data models.UserPackage

	err := e.Orm.Model(&data).Debug().
		Scopes(
			cDto.MakeCondition(d.GetNeedSearch()),
			cDto.Paginate(d.GetPageSize(), d.GetPageIndex()),
		).
		Find(list, "user_id = ?", d.GetUserId()).
		Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// GetPackagePagesByIds 获取package列表
func (e *UserPackage) GetPackagePagesByIds(d *dto.PackagesByIdsForRelationshipUsers, list *[]models.Package, count *int64) error {
	var err error
	var ids []int = d.GetPackageId()
	for i := 0; i < len(ids); i++ {
		if ids[i] != 0 {
			var data1 models.Package
			err = e.Orm.Model(&data1).
				Where("package_id = ? ", ids[i]).
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

// InsertUserPackage 创建
func (e Package) InsertUserPackage(c *dto.UserPackageInsertReq) error {
	var err error
	var data models.UserPackage
	var i int64
	fmt.Println("rpid，ruid已查出:", c.PackageId, c.UserId)
	fmt.Println(e.Orm)
	err = e.Orm.Model(&data).Where("user_id = ? and package_id = ?", c.UserId, c.PackageId).Error
	fmt.Println("3pid，uid已查出:")
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("关系已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}

	c.GenerateUserPackage(&data)
	fmt.Println("4pid，uid已查出:")
	err = e.Orm.Create(&data).Error
	fmt.Println("5pid，uid已查出:", data.Id)
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// RemoveRelationship 根据专利包id删除用户专利包关系
func (e *UserPackage) RemoveRelationship(c *dto.UserPackageObject) error {
	var err error
	var data models.UserPackage

	db := e.Orm.Where("package_id = ? AND user_id = ? ", c.PackageId, c.UserId).
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
