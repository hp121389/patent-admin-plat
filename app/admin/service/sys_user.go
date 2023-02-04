package service

import (
	"errors"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"gorm.io/gorm"

	cDto "go-admin/common/dto"
)

type SysUser struct {
	service.Service
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(c *dto.SysUserGetPageReq, list *[]models.SysUser, count *int64) error {
	var err error
	//var data models.SysUser
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

// Get 获取SysUser对象
func (e *SysUser) Get(d *dto.SysUserById, model *models.SysUser) error {
	var data models.SysUser

	err := e.Orm.Model(&data).Debug().
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

// Insert 创建SysUser对象
func (e *SysUser) Insert(c *dto.SysUserInsertReq) error {
	var err error
	var data models.SysUser
	var i int64
	err = e.Orm.Model(&data).Where("username = ?", c.Username).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("用户名已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	if data.DepartID != 0 {
		deptRela := models.DeptRelation{
			UserId: data.UserId,
			DeptId: data.DepartID,
		}
		err = e.Orm.Create(&deptRela).Error
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}
	}

	return nil
}

// Update 修改SysUser对象
func (e *SysUser) Update(c *dto.SysUserUpdateReq) error {
	var err error
	var data models.SysUser
	db := e.Orm.First(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&data)
	update := e.Orm.Model(&data).Where("user_id = ?", &data.UserId).Omit("password", "salt").Updates(&data)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update userinfo error")
		log.Warnf("db update error")
		return err
	}

	if c.DepartID != 0 {
		deptRela := models.DeptRelation{
			UserId: data.UserId,
			DeptId: data.DepartID,
		}
		err = e.Orm.Model(&deptRela).
			Where("user_id = ?", data.UserId).
			Updates(&deptRela).
			Error
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}
	}

	return nil
}

// InternalUpdate internal修改SysUser对象
func (e *SysUser) InternalUpdate(c *dto.SysUserUpdateReq) error {
	var err error
	var data models.SysUser
	db := e.Orm.First(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	c.Generate(&data)
	update := e.Orm.Model(&data).Where("user_id = ?", &data.UserId).Omit("password", "salt").Updates(&data)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update userinfo error")
		log.Warnf("db update error")
		return err
	}

	if c.DepartID != 0 {
		deptRela := models.DeptRelation{
			UserId: data.UserId,
			DeptId: data.DepartID,
		}
		db = e.Orm.Model(&deptRela).
			Where("user_id = ?", data.UserId).
			Updates(&deptRela)
		if db.Error != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}
		if db.RowsAffected == 0 {
			if err = e.Orm.Create(&deptRela).Error; err != nil {
				e.Log.Errorf("db error: %s", err)
				return err
			}
		}
	}

	return nil
}

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(c *dto.UpdateSysUserAvatarReq) error {
	var err error
	var model models.SysUser
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = e.Orm.Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(c).Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	return nil
}

// UpdateStatus 更新用户状态
func (e *SysUser) UpdateStatus(c *dto.UpdateSysUserStatusReq) error {
	var err error
	var model models.SysUser
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = e.Orm.Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(c).Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	return nil
}

// ResetPwd 重置用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq) error {
	var err error
	var model models.SysUser
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("At Service ResetSysUserPwd error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	c.Generate(&model)
	err = e.Orm.Omit("username", "nick_name", "phone", "role_id", "avatar", "sex").Save(&model).Error
	if err != nil {
		e.Log.Errorf("At Service ResetSysUserPwd error: %s", err)
		return err
	}
	return nil
}

// Remove 删除SysUser
func (e *SysUser) Remove(c *dto.SysUserById) error {
	var err error
	var data models.SysUser

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

// UpdatePwd 修改SysUser对象密码
func (e *SysUser) UpdatePwd(id int, oldPassword, newPassword string) error {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err = e.Orm.Model(c).
		Select("UserId", "Password", "Salt").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("无权更新该数据")
		}
		e.Log.Errorf("db error: %s", err)
		return err
	}
	var ok bool
	ok, err = pkg.CompareHashAndPassword(c.Password, oldPassword)
	if err != nil {
		e.Log.Errorf("CompareHashAndPassword error, %s", err.Error())
		return err
	}
	if !ok {
		err = errors.New("incorrect Password")
		e.Log.Warnf("user[%d] %s", id, err.Error())
		return err
	}
	c.Password = newPassword
	db := e.Orm.Model(c).Where("user_id = ?", id).
		Select("Password", "Salt").
		Updates(c)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("set password error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

func (e *SysUser) GetProfile(c *dto.SysUserById, user *models.SysUser, roles *[]models.SysRole) error {
	err := e.Orm.Preload("SysDept").First(user, c.GetId()).Error
	if err != nil {
		return err
	}
	err = e.Orm.Find(roles, user.RoleId).Error
	if err != nil {
		return err
	}

	return nil
}
