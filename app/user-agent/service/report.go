package service

import (
	"errors"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/google/uuid"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service/dtos"
	"gorm.io/gorm"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	//"gorm.io/gorm"
)

type Report struct {
	service.Service
}

// UserGetRepoById 获取Report对象
func (e *Report) UserGetRepoById(id int, model *model.Report) error {
	//引用传递、函数名、形参、返回值
	var err error
	db := e.Orm.Where("report_Id = ?  ", id).First(model)
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

// InsertReport 添加Report对象
func (e *Report) InsertReport(c *dtos.ReportGetPageReq, typer string) (error, *model.Report) {

	var err error
	var data model.Report
	var i int64
	err = e.Orm.Model(&data).Where("report_id = ?", c.ReportId).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err, nil
	}
	if i > 0 {
		err := errors.New("报告ID已存在！")
		e.Log.Errorf("db error: %s", err)
		return err, nil
	}

	c.Generate(&data)
	data.CreatedAt = dtos.UpdateTime()
	data.ReportName = typer + ".defaultName." + uuid.New().String()
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err, nil
	}

	return nil, &data
}

func (e *Report) InsertRelation(c *dtos.ReportRelaReq) error {
	var err error
	var data model.ReportRelation
	var i int64

	err = e.Orm.Model(&data).
		Where("report_id = ? and user_id = ? and patent_id = ?", c.ReportId, c.UserId, c.PatentId).
		Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("该关系已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.GenerateRela(&data)
	data.CreatedAt = dtos.UpdateTime()

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

// UpdateReport 撤销申请报告
func (e *Report) UpdateReport(c *dtos.ReportGetPageReq) error {
	var err error
	var model model.Report
	db := e.Orm.Where("Report_Id = ? ", c.ReportId).First(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update report error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("报告不存在")
	}
	c.Generate(&model)
	model.UpdatedAt = dtos.UpdateTime()
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

//
//// DeleteReportsRela 删除申请关系
//func (e *Report) DeleteReportsRela(c *dtos.ReportRelaReq) error {
//	var err error
//	var model model.ReportRelation
//	db := e.Orm.Where("Report_Id = ?  ", c.ReportId).First(&model)
//	if err = db.Error; err != nil {
//		e.Log.Errorf("Service Update report error: %s", err)
//		return err
//	}
//	if db.RowsAffected == 0 {
//		return errors.New("报告不存在")
//
//	}
//	c.GenerateRela(&model)
//	model.UpdatedAt = dtos.UpdateTime()
//
//	update := e.Orm.Model(&model).Where("report_id = ?", &model.ReportId).Delete(&model)
//	if err = update.Error; err != nil {
//		e.Log.Errorf("db error: %s", err)
//		return err
//	}
//	if update.RowsAffected == 0 {
//		err = errors.New("update report-info error")
//		log.Warnf("db update error")
//		return err
//	}
//
//	return nil
//}
