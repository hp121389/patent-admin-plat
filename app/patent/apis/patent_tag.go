package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/patent/models"
	"go-admin/app/patent/service"
	"go-admin/app/patent/service/dto"
)

type PatentTag struct {
	api.Api
}

// GetPatentListsByUserId
// @Summary 列表专利信息数据
// @Description 获取JSON
// @Tags 专利表
// @Router /api/v1/patent-list/get_patent_lists [get]
// @Security Bearer
func (e PatentTag) GetPatentListsByUserId(c *gin.Context) { //gin框架里的上下文
	s := service.PatentTag{}         //service中查询或者返回的结果赋值给s变量
	req := dto.PatentTagGetPageReq{} //被绑定的数据
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

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	list := make([]models.PatentTag, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}
