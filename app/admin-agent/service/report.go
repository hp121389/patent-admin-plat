package service

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/service/dto"
	"gorm.io/gorm"
)

type Report struct {
	service.Service
}

// GetValuationPatentById 获取专利-估值对象
func (e *Report) GetValuationPatentById(d *dto.PatentReq, model *model.PatentReport) error {

	var err error
	db := e.Orm.Where("patent_Id = ? ", d.PatentId).First(model)
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

// GetValuationReportById 获取ValuationReport对象
func (e *Report) GetValuationReportById(d *dtos.ReportById, model *model.Report) error {
	//引用传递、函数名、形参、返回值
	var err error
	db := e.Orm.Where("report_Id = ? and Type = ?", d.ReportId, dtos.InfringementType).First(model)
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
func (e *Report) GetInfringementReportById(d *dtos.ReportById, model *model.Report) error {
	//引用传递、函数名、形参、返回值
	var err error
	fmt.Println(d.ReportId)
	db := e.Orm.Where("report_Id = ? and Type = ?", d.ReportId, dtos.InfringementType).First(model)
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
	var data []model.Report

	err = e.Orm.Model(&data).Where("Type = ?", dtos.ValuationType).
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
	var data []model.Report

	err = e.Orm.Model(&data).Where("Type = ?", dtos.InfringementType).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetPatentByReId 获取Report对象
func (e *Report) GetPatentByReId(reid int, model *model.PatentReport) error {

	var err error
	//fmt.Println("before")
	//fmt.Println(reid)
	db := e.Orm.Where("report_Id = ? ", reid).First(model)
	//用First(model，reid)会有不一样的结果，但是也是错误的
	//fmt.Println("after")
	//fmt.Println("model.ReportId")
	//fmt.Println(model.ReportId)
	//fmt.Println("model.PatentId") //1：为什么一直是1
	//fmt.Println(model.PatentId)   //1：为什么一直是1:First在where后
	//fmt.Println("model.Id")
	//fmt.Println(model.Id)
	//fmt.Println("走出service")
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("报告对应专利不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}
