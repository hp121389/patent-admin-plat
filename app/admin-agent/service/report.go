package service

import (
	"errors"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/service/dto"
	"gorm.io/gorm"
	"time"
)

type Report struct {
	service.Service
}

//-----------------------------------------Get-----------------------------------------------------------

// GetReportById 获取ValuationReport对象
func (e *Report) GetReportById(d *dtos.ReportById, model *model.Report) error {
	//引用传递、函数名、形参、返回值
	var err error
	db := e.Orm.Where("report_Id = ?  ", d.ReportId).First(model)
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

//--------------------------------------------patent---------------------------------------------------------

// GetPatentByReId 获取Report对象
func (e *Report) GetPatentByReId(reid int, model *model.PatentReport) error {

	var err error
	db := e.Orm.Where("report_Id = ? ", reid).First(model)

	//fmt.Println(model.PatentId)   //1：为什么一直是1:First在where后

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

// UpdateReports 更新报告对象
func (e *Report) UpdateReports(c *dtos.ReportGetPageReq) error {
	var err error
	var model model.Report
	db := e.Orm.Where("Report_Id = ?  ", c.ReportId).First(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update report error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("专利不存在")

	}
	c.GenerateNoneFile(&model)
	CurrentTime1 := fmt.Sprintf("%v", time.Now())
	CurrentTime2 := CurrentTime1[0:19]
	model.UpdatedAt = CurrentTime2
	update := e.Orm.Model(&model).Where("report_id = ?", &model.ReportId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update report-info error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// UploadReport 上传报告
func (e *Report) UploadReport(c *dtos.ReportGetPageReq) error {
	var err error
	var model model.Report
	db := e.Orm.First(&model, c.ReportId)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("package not found")

	}

	switch c.FilesOpt {
	case dto.FilesAdd:
		c.GenerateAndAddFiles(&model)
	case dto.FilesDelete:
		c.GenerateAndDeleteFiles(&model)
	default:
		c.Generate(&model)
	}
	CurrentTime1 := fmt.Sprintf("%v", time.Now())
	CurrentTime2 := CurrentTime1[0:19]
	model.UpdatedAt = CurrentTime2
	update := e.Orm.Model(&model).Where("report_id = ?", &model.ReportId).Updates(&model)
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
