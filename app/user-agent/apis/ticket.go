package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service"
	"go-admin/app/admin-agent/service/dtos"
	serviceUser "go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"strconv"
)

type Ticket struct {
	api.Api
}

// GetTicketPages
// @Summary 获取工单列表
// @Description 获取工单列表
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/tickets [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param type query string true "type"
// @Param query query string true "query"
// @Param status query string true "status"
// @Param reportType query string false "reportType"
// @Security Bearer
func (e Ticket) GetTicketPages(c *gin.Context) {
	s := service.Ticket{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	t := c.Query("type")
	status := c.Query("status")
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dtos.TicketPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Type = t
	req.Status = status
	req.UserID = user.GetUserId(c)
	req.Query = c.Query("query")

	list := make([]model.Ticket, 0)
	var count int64
	if err = s.GetTicketPages(&req, &list, &count); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	switch t {
	case dtos.TicketTypeReport:
		rs := service.Report{}
		err = e.MakeContext(c).
			MakeOrm().
			MakeService(&rs.Service).
			Errors
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}

		rt := c.Query("reportType")
		var _c int64
		list, err = rs.GetReportTicketListByTickets(rt, list, &_c)
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
	}

	e.PageOK(list, int(count), req.PageSize, req.PageIndex, "查询成功")
}

// CreateTicket
// @Summary 新建工单
// @Description 新建工单
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Param data body dtos.TicketDBReq true "报告工单参数"
// @Router /api/v1/user-agent/tickets [post]
// @Security Bearer
func (e Ticket) CreateTicket(c *gin.Context) {
	ticketType := c.Query("type")
	ticketDBReq := dtos.TicketDBReq{}
	var ticket *model.Ticket
	var err error
	switch ticketType {
	case dtos.TicketTypeReport:
		req := dtos.NewReportTicketReq()
		rs := service.Report{}
		err = e.MakeContext(c).
			MakeOrm().
			Bind(&req).
			MakeService(&rs.Service).
			Errors
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		userID := user.GetUserId(c)
		req.UserID = userID

		var report *model.Report
		if report, err = rs.Create(&req.RelObj); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		ticketDBReq = dtos.TicketDBReq{
			RelaID:     report.ReportId,
			Name:       fmt.Sprintf("报告工单：%s", req.RelObj.ReportName),
			Properties: req.Properties,
			Type:       dtos.TicketTypeReport,
			UserID:     req.UserID,
			OptMsg:     fmt.Sprintf("新建工单：申请%s报告%s", req.RelObj.Type, req.RelObj.ReportName),
		}

		// do something
		defer func() {
			if ticket != nil {
				rrr := dtos.ReportRelaReq{
					TicketId: ticket.ID,
					ReportId: report.ReportId,
					UserId:   userID,
				}
				if err = rs.Link(&rrr); err != nil {
					e.Logger.Error(err)
					e.Error(500, err, err.Error())
					return
				}

				// novelty report
				if req.RelObj.Type == dtos.ReportTypeNovelty {
					noveltyReq, err := convertToNoveltyReq(req.Properties)
					if err != nil {
						e.Logger.Error(err)
						e.Error(500, err, err.Error())
						return
					}
					go e.genPatentNovelty(&rrr, noveltyReq)
				}
			}
		}()
	case dtos.TicketTypeCommon:
		req := dtos.TicketDBReq{}
		err = e.MakeContext(c).
			Bind(&req).
			Errors
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		req.Type = dtos.TicketTypeCommon
		req.OptMsg = fmt.Sprintf("新建工单：%s", req.Name)
		ticketDBReq = req
	default:
		err = fmt.Errorf("invalid ticket type: %s", ticketType)
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	s := service.Ticket{}
	err = e.MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	if ticket, err = s.Create(&ticketDBReq); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(nil, "创建成功")
}

// UpdateTicket
// @Summary 更新工单
// @Description 更新工单
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/tickets/{id} [put]
// @Security Bearer
func (e Ticket) UpdateTicket(c *gin.Context) {
	ticketType := c.Query("type")
	ticketDBReq := dtos.TicketDBReq{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	switch ticketType {
	case dtos.TicketTypeReport:
		req := dtos.NewReportTicketReq()
		rs := service.Report{}
		err = e.MakeContext(c).
			MakeOrm().
			Bind(&req).
			MakeService(&rs.Service).
			Errors
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}

		relaReq := dtos.ReportRelaReq{TicketId: id}
		rela, err := rs.GetReportRelaByTicketId(&relaReq)
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}

		req.RelObj.ReportId = rela.ReportId
		if err = rs.Update(&req.RelObj); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}

		req.GenOptLogsWhenUpdate()
		ticketDBReq = req.TicketDBReq
	case dtos.TicketTypeCommon:
		req := dtos.TicketDBReq{}
		err = e.MakeContext(c).
			Bind(&req).
			Errors
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}

		req.GenOptLogsWhenUpdate()
		ticketDBReq = req
	default:
		err = fmt.Errorf("invalid ticket type: %s", ticketType)
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)
	ticketDBReq.UserID = userID

	s := service.Ticket{}
	err = e.MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	if err = s.Update(id, &ticketDBReq); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(nil, "更新成功")
}

// CloseTicket
// @Summary 关闭工单
// @Description 关闭工单
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/tickets/{id}/close [put]
// @Security Bearer
func (e Ticket) CloseTicket(c *gin.Context) {
	ts := service.Ticket{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&ts.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	tid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	if err = ts.Close(tid, dtos.NewTicketDBReq(user.GetUserId(c), "用户侧关闭工单")); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(nil, "更新成功")
}

func (e Ticket) genPatentNovelty(rrr *dtos.ReportRelaReq, req *dto.NoveltyReportReq) {
	var reportResp *dto.NoveltyReportResp
	var err error

	s := serviceUser.Report{}
	err = e.MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	ars := service.Report{}
	err = e.MakeOrm().
		MakeService(&ars.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	ts := service.Ticket{}
	err = e.MakeOrm().
		MakeService(&ts.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	defer func() {
		if err != nil {
			if err = ts.Update(rrr.TicketId,
				dtos.NewTicketDBReq(rrr.UserId, fmt.Sprintf("查新报告生成失败，失败原因: %s", err))); err != nil {
				e.Logger.Error(err)
				return
			}
		}
	}()

	err = ts.Update(rrr.TicketId, dtos.NewTicketDBReq(rrr.UserId, "系统自动生成查新报告中..."))
	if err != nil {
		e.Logger.Error(err)
		return
	}

	reportResp, err = s.GetNovelty(req)
	if err != nil {
		e.Logger.Error(err)
		return
	}

	reportUpdate := dtos.ReportReq{
		ReportId:         rrr.ReportId,
		ReportProperties: reportResp.Map(),
	}
	if err = ars.Update(&reportUpdate); err != nil {
		e.Logger.Error(err)
		return
	}

	if err = ts.Update(rrr.TicketId, dtos.NewTicketDBReq(rrr.UserId, "查新报告生成成功")); err != nil {
		e.Logger.Error(err)
		return
	}
	if err = ts.Finish(rrr.TicketId, dtos.NewTicketDBReq(rrr.UserId, "报告生成结束，自动关闭")); err != nil {
		e.Logger.Error(err)
		return
	}
}

func convertToNoveltyReq(properties dtos.Properties) (*dto.NoveltyReportReq, error) {
	src := []byte(properties.String())
	noveltyReq := dto.NoveltyReportReq{}
	if err := json.Unmarshal(src, &noveltyReq); err != nil {
		return nil, err
	}
	return &noveltyReq, nil
}
