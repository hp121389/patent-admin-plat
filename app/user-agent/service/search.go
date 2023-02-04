package service

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
	cDto "go-admin/common/dto"
)

type Search struct {
	service.Service
}

func (e *Search) GetQueryPage(c *dto.StoredQueryReq, list *[]models.StoredQuery, count *int64) error {
	var err error
	var data models.StoredQuery
	err = e.Orm.Model(&data).
		Scopes(cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *Search) FindQueryPages(c *dto.StoredQueryFindReq, list *[]models.StoredQuery, count *int64) error {
	var err error
	var data models.StoredQuery
	err = e.Orm.Model(&data).
		Scopes(cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", c.Query)).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *Search) InsertQuery(c *dto.StoredQueryReq) error {
	var err error
	var data models.StoredQuery
	var i int64
	err = e.Orm.Model(&data).Where("name = ?", c.Name).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err = fmt.Errorf("storedQuery name %s has existed", c.Name)
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.QueryID = data.QueryID
	return nil
}

func (e *Search) RemoveQuery(id int) error {
	var err error
	var data models.StoredQuery

	db := e.Orm.Model(&data).
		Delete(&data, id)
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveQuery : %s", err)
		return err
	}
	return nil
}

func (e *Search) UpdateQuery(id int, c *dto.StoredQueryReq) error {
	var err error
	var model models.StoredQuery
	db := e.Orm.First(&model, id)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateQuery error: %s", err)
		return err
	}

	if db.RowsAffected == 0 {
		return errors.New("storedQuery is not existed")
	}

	c.Generate(&model)
	update := e.Orm.Model(&model).Where("query_id = ?", id).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update stored query error")
		e.Log.Errorf("db update error", err)
		return err
	}
	return nil
}
