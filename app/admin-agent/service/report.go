package service

import (
	"errors"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/service/dto"
	"gorm.io/gorm"
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

// GetPagesByType 获取Type类型Report对象列表
func (e *Report) GetPagesByType(typeRepo string, list *[]model.Report) error {
	var err error
	var data []model.Report
	err = e.Orm.Model(&data).Where("Type = ?", typeRepo).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

//--------------------------------------------patent---------------------------------------------------------

// GetPatentByReId 获取patent对象
func (e *Report) GetPatentByReId(reid int, model *model.ReportRelation) error {

	var err error
	db := e.Orm.Where("report_Id = ? ", reid).First(model)
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

// GetReportIdsByPatentId 通过PatentId获取report对象ids
func (e *Report) GetReportIdsByPatentId(patentId int, userId int, list *[]model.ReportRelation) error {

	var err error
	db := e.Orm.Where("patent_id = ? and user_id = ?", patentId, userId).Find(list)
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

// GetReportIdsByUserId 通过UserId获取report对象ids
func (e *Report) GetReportIdsByUserId(userId int, list *[]model.ReportRelation) error {

	var err error
	db := e.Orm.Where("user_id = ? ", userId).Find(list)
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

// GetReportIdsByType 通过type获取report对象ids
func (e *Report) GetReportIdsByType(d *dtos.ReportRelaReq, list *[]model.ReportRelation) error {
	var err error
	db := e.Orm.Where("user_id = ? and type = ?", d.UserId, d.Type).Find(list)
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

// GetReportListByIds 通过ReportIds找到对应的ReportList
func (e *Report) GetReportListByIds(d *dtos.ReportIds, list *[]model.Report) error {
	var err error
	var ids []int = d.GetReportId()
	for i := 0; i < len(ids); i++ {
		if ids[i] != 0 {
			var data1 model.Report
			err = e.Orm.Model(&data1).
				Where("report_Id = ? ", ids[i]).
				First(&data1).Limit(-1).Offset(-1).
				Error
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

// UploadReport 上传报告
func (e *Report) UploadReport(c *dtos.ReportGetPageReq) error {
	var err error
	var model model.Report
	db := e.Orm.First(&model, c.ReportId)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Report error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("report not found")
	}

	switch c.FilesOpt {
	case dto.FilesAdd:
		c.GenerateAndAddFiles(&model)
	case dto.FilesDelete:
		c.GenerateAndDeleteFiles(&model)
	default:
		c.Generate(&model)
	}

	model.UpdatedAt = dtos.UpdateTime()
	update := e.Orm.Model(&model).Where("report_id = ?", &model.ReportId).Updates(&model)

	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update report error")
		log.Warnf("db update error")
		return err
	}
	return nil
}
