package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service/dto"
	"gorm.io/gorm"
)

type Report struct {
	service.Service
}

// GetValuationReportById 获取ValuationReport对象
func (e *Report) GetValuationReportById(d *dto.ReportById, model *model.Report) error {
	//引用传递、函数名、形参、返回值
	var err error
	db := e.Orm.First(model).Where("report_Id = ? and Type = ?", d.ReportId, dto.ValuationType)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看估值报告不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetInfringementReportById 获取InfringementReport对象
func (e *Report) GetInfringementReportById(d *dto.ReportById, model *model.Report) error {
	//引用传递、函数名、形参、返回值
	var err error
	db := e.Orm.First(model).Where("report_Id = ? and Type = ?", d.ReportId, dto.InfringementType)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看侵权报告不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetValuationReportPages 获取ValuationReport对象列表
func (e *Report) GetValuationReportPages(list *[]model.Report) error {
	var err error
	var data model.Report

	err = e.Orm.Model(&data).Where("Type = ?", dto.ValuationType).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetInfringementReportPages 获取InfringementReport对象列表
func (e *Report) GetInfringementReportPages(list *[]model.Report) error {
	var err error
	var data model.Report

	err = e.Orm.Model(&data).Where("Type = ?", dto.InfringementType).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}
