package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/patent/service"
	"go-admin/app/patent/service/dto"
)

type UserTag struct {
	api.Api
}

// Insert
// @Summary 增加标签
// @Description 获取JSON
// @Tags 标签--用户/Tag-User
// @Param data body dto.UserTagInsertReq true "标签数据"
// @Router /api/v1/tag [post]
// @Security Bearer
func (e UserTag) Insert(c *gin.Context) {
	s := service.UserTag{}
	req := dto.UserTagInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetId(), "创建成功")
}
