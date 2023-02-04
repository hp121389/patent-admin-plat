package service

import (
	"errors"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/service/dto"
	cDto "go-admin/common/dto"
)

type Report struct {
	service.Service
}

func (e *Report) Create(req *dtos.ReportReq) (*model.Report, error) {
	var err error
	var data model.Report
	var i int64
	err = e.Orm.Model(&data).Where("report_name = ?", req.ReportName).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if i > 0 {
		err = fmt.Errorf("report with report_name=%s existed", req.ReportName)
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	req.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	return &data, nil
}

func (e *Report) Link(req *dtos.ReportRelaReq) error {
	var err error
	var data model.ReportRelation

	req.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *Report) GetReportRelaByTicketId(req *dtos.ReportRelaReq) (*model.ReportRelation, error) {
	var err error
	var data model.ReportRelation

	req.Generate(&data)
	err = e.Orm.Model(&data).
		Where(&data).
		First(&data).
		Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	return &data, nil
}

func (e *Report) GetReportPagesByTickets(req *dtos.ReportPagesReq, tickets []model.Ticket, count *int64) ([]model.Report, error) {
	if len(tickets) == 0 {
		*count = 0
		return []model.Report{}, nil
	}

	ticketIDs := make([]int, 0, len(tickets))
	ticketsMap := make(map[int]model.Ticket)
	for _, t := range tickets {
		ticketIDs = append(ticketIDs, t.ID)
		ticketsMap[t.ID] = t
	}

	tl := model.ReportRelation{}
	tl.UserId = req.UserID
	tls := make([]model.ReportRelation, 0)
	err := e.Orm.Model(&tl).
		Where(&tl).
		Where("ticket_id IN ?", ticketIDs).
		Find(&tls).Limit(-1).Offset(-1).
		Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}

	reportIDs := make([]int, 0, len(tickets))
	tlMap := make(map[int]model.ReportRelation)
	for _, rel := range tls {
		reportIDs = append(reportIDs, rel.ReportId)
		tlMap[rel.ReportId] = rel
	}

	reports := make([]model.Report, 0)
	var data model.Report
	data.Type = req.Type
	err = e.Orm.Model(&data).
		Scopes(cDto.Paginate(req.GetPageSize(), req.GetPageIndex())).
		Where(&data).
		Where("report_name LIKE ?", fmt.Sprintf("%%%s%%", req.Query)).
		Find(&reports, reportIDs).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}

	return reports, nil
}

func (e *Report) GetReportTicketListByTickets(reportType string, tickets []model.Ticket, count *int64) ([]model.Ticket, error) {
	ticketIDs := make([]int, 0, len(tickets))
	ticketsMap := make(map[int]model.Ticket)
	for _, t := range tickets {
		ticketIDs = append(ticketIDs, t.ID)
		ticketsMap[t.ID] = t
	}

	tl := model.ReportRelation{}
	tls := make([]model.ReportRelation, 0)
	err := e.Orm.Model(&tl).
		Where("ticket_id IN ?", ticketIDs).
		Find(&tls).Limit(-1).Offset(-1).
		Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}

	reportIDs := make([]int, 0, len(tickets))
	tlMap := make(map[int]model.ReportRelation)
	for _, rel := range tls {
		reportIDs = append(reportIDs, rel.ReportId)
		tlMap[rel.ReportId] = rel
	}

	reports := make([]model.Report, 0)
	var data model.Report
	data.Type = reportType
	err = e.Orm.Model(&data).
		Where(&data).
		Find(&reports, reportIDs).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return nil, err
	}

	res := make([]model.Ticket, 0, len(tickets))
	for _, report := range reports {
		tid := tlMap[report.ReportId].TicketId
		if ticket, ok := ticketsMap[tid]; ok {
			resTicket := ticket
			resTicket.RelObj = report
			res = append(res, resTicket)
		}
	}
	return res, nil
}

func (e *Report) Update(c *dtos.ReportReq) error {
	var err error
	var r model.Report
	db := e.Orm.First(&r, c.ReportId)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateReport error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("report not found")
	}

	switch c.FilesOpt {
	case dto.FilesAdd:
		c.GenerateAndAddFiles(&r)
	case dto.FilesDelete:
		c.GenerateAndDeleteFiles(&r)
	default:
		c.Generate(&r)
	}

	update := e.Orm.Model(&r).Where("report_id = ?", &r.ReportId).Updates(&r)
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
