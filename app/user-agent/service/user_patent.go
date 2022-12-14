package service

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	models "go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
)

type UserPatent struct {
	service.Service
}

// GetUserPatentIds 通过UserId获得专利列表的ID数组
func (e *UserPatent) GetUserPatentIds(c *dto.UserPatentObject, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("User_Id = ?", c.UserId).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

//// GetUserFocusPatentIds 通过UserId获得专利列表的ID数组
//func (e *UserPatent) GetUserFocusPatentIds(c *dto.UserPatentGetPageReq, list *[]models.UserPatent) error {
//	var err error
//	var data models.UserPatent
//	fmt.Println(c)
//
//	err = e.Orm.Model(&data).
//		Where("User_Id = ? AND type = ?", c.GetUserId(), dto.FocusType).
//		Find(&list).Limit(-1).Offset(-1).Error
//	fmt.Println(list)
//	if err != nil {
//		e.Log.Errorf("db error:%s", err)
//		return err
//	}
//	return nil
//}

// GetClaimLists 通过专利列表的ID数组获得认领专利列表
func (e *UserPatent) GetClaimLists(c *dto.UserPatentObject, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("Type = ? AND User_Id = ?", dto.ClaimType, c.UserId).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	return nil
}

// GetFocusLists 通过专利列表的ID数组获得关注专利列表
func (e *UserPatent) GetFocusLists(c *dto.UserPatentObject, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("Type = ? AND User_Id = ?", dto.FocusType, c.UserId).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetAllRelatedPatentsByUserId 通过专利列表的ID数组获得与该用户相关的所有(认领+关注)专利列表
func (e *UserPatent) GetAllRelatedPatentsByUserId(d *dto.UserPatentObject, list *[]models.UserPatent) error {
	var err error
	err = e.Orm.Debug().
		Where("user_id = ?", d.UserId).
		Find(list).Limit(-1).Offset(-1).
		Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// RemoveClaim 取消认领
func (e *UserPatent) RemoveClaim(c *dto.UserPatentObject) error {
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
func (e *UserPatent) RemoveFocus(c *dto.UserPatentObject) error {
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

// InsertUserPatent insert relationship between user and patent
func (e *UserPatent) InsertUserPatent(c *dto.UserPatentObject) error {
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

// GetUsersByPatentId  通过专利id数组获取关注或者认领的用户
func (e *UserPatent) GetUsersByPatentId(list *[]models.UserPatent, pid *dto.PatentsIds) error {
	var err error
	var data models.UserPatent
	//result := make([]models.UserPatent,0)
	fmt.Println(pid.PatentIds)
	for i := 0; i < len(pid.PatentIds); i++ {
		var templist []models.UserPatent
		//fmt.Print("now patent is:")
		//fmt.Println(pid.PatentIds[i])
		err = e.Orm.Where(&data).Where("Patent_Id = ? ", pid.PatentIds[i]).Find(&templist).Limit(100000).Error
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}
		//fmt.Print("now user-patent is:")
		fmt.Println(len(templist))
		for j := 0; j < len(templist); j++ {
			//fmt.Println(templist[j])
			*list = append(*list, templist[j])
		}
	}
	return nil
}

// GetUsersByPatentId  通过专利id数组获取关注或者认领的用户
func (e *UserPatent) GetFocusUsersByPatentId(list *[]models.UserPatent, pid *dto.PatentsIds) error {
	var err error
	var data models.UserPatent
	//result := make([]models.UserPatent,0)
	fmt.Println(pid.PatentIds)
	for i := 0; i < len(pid.PatentIds); i++ {
		var templist []models.UserPatent
		//fmt.Print("now patent is:")
		//fmt.Println(pid.PatentIds[i])
		err = e.Orm.Where(&data).Where("Patent_Id = ? and type = ? ", pid.PatentIds[i], dto.FocusType).Find(&templist).Limit(100000).Error
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}
		//fmt.Print("now user-patent is:")
		fmt.Println(len(templist))
		for j := 0; j < len(templist); j++ {
			//fmt.Println(templist[j])
			*list = append(*list, templist[j])
		}
	}
	return nil
}

//// Get
//func (e *UserPatent) GetTwoUserRelationshipInThisPackage(plist *dto.PatentsIds, member1 int, member2 int) (int, error) {
//	var data models.UserPatent
//	sum := 0
//	for i := 0; i < len(plist.PatentIds); i++ {
//		var count1 int64
//		var count2 int64
//		err := e.Orm.Model(data).Where("Patent_Id = ? and User_Id = ?", plist.PatentIds[i], member1).Count(&count1).Error
//		err = e.Orm.Model(data).Where("Patent_Id = ? and User_Id = ?", plist.PatentIds[i], member2).Count(&count2).Error
//		if err != nil {
//			e.Log.Errorf("db error: %s", err)
//			return 0, err
//		}
//		if count1 != 0 && count2 != 0 {
//			sum++
//		}
//
//	}
//	return sum, nil
//
//}
