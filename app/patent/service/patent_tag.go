package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/patent/models"
	"go-admin/app/patent/service/dto"
	cDto "go-admin/common/dto"
)

type PatentTag struct {
	service.Service
}

// GetPage 获取patent列表
func (e *PatentTag) GetPage(c *dto.PatentTagGetPageReq, list *[]models.PatentTag, count *int64) error {
	var err error
	var data models.PatentTag

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
