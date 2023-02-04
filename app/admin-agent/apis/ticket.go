package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service"
	"go-admin/app/admin-agent/service/dtos"
	"strconv"
)

type Ticket struct {
	api.Api
}

// GetAllTicketPages
// @Summary 获取工单列表
// @Description 获取工单列表
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/tickets [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param type query string true "type"
// @Param query query string true "type"
// @Param status query string true "status"
// @Param reportType query string false "reportType"
// @Security Bearer
func (e Ticket) GetAllTicketPages(c *gin.Context) {
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

		var _c int64
		rt := c.Query("reportType")
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
// @Param data body dtos.TicketDBReq true "工单数据"
// @Router /api/v1/admin-agent/tickets [post]
// @Security Bearer
func (e Ticket) CreateTicket(c *gin.Context) {
	s := service.Ticket{}
	req := dtos.TicketDBReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)
	req.UserID = userID

	if _, err = s.Create(&req); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "创建成功")
}

// UpdateTicket
// @Summary 更新工单
// @Description 更新工单
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/tickets/{id} [put]
// @Param type query string true "type"
// @Param data body dtos.TicketDBReq true "工单数据"
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

		req.GenOptLogsWhenUpdate()
		ticketDBReq = req.TicketDBReq
		if err = rs.Update(&req.RelObj); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
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
// @Router /api/v1/admin-agent/tickets/{id}/close [put]
// @Security Bearer
func (e Ticket) CloseTicket(c *gin.Context) {
	s := service.Ticket{}
	req := dtos.TicketDBReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)
	req.UserID = userID

	if err = s.Close(id, &req); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "关闭成功")
}

// FinishTicket
// @Summary 完结工单
// @Description 完结工单
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/tickets/{id}/finish [put]
// @Security Bearer
func (e Ticket) FinishTicket(c *gin.Context) {
	s := service.Ticket{}
	req := dtos.TicketDBReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)
	req.UserID = userID

	if err = s.Finish(id, &req); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "关闭成功")
}

// RemoveTicket
// @Summary 删除工单
// @Description 删除工单
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/tickets/{id} [delete]
// @Security Bearer
func (e Ticket) RemoveTicket(c *gin.Context) {
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	if err = s.Remove(id); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(nil, "删除成功")
}
