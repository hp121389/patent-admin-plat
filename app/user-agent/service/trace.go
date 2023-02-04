package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
	cDto "go-admin/common/dto"
)

type Tracer struct {
	service.Service
}

func (e *Tracer) Trace(req *dto.TraceReq) error {
	var err error
	var data models.TraceLog
	req.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *Tracer) SelectTraceLog(req *dto.TracePageReq, list *[]models.TraceLog, count *int64) error {
	var err error
	switch {
	case len(req.Action) != 0:
		err = e.Orm.Debug().
			Where("action = ? AND create_by = ?", req.Action, req.UserID).
			Scopes(
				cDto.Paginate(req.GetPageSize(), req.GetPageIndex()),
			).
			Find(list).Limit(-1).Offset(-1).
			Count(count).Error
	default:
		err = e.Orm.Debug().
			Where("create_by = ?", req.UserID).
			Scopes(
				cDto.Paginate(req.GetPageSize(), req.GetPageIndex()),
			).
			Find(list).Limit(-1).Offset(-1).
			Count(count).Error
	}

	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
