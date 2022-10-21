package dto

import (
	"go-admin/app/user-agent/models"
	common "go-admin/common/models"
)

type UserTagInsertReq struct {
	UserTagId int `json:"UserTagId" comment:"用户和标签绑定的id"` //用户-标签关系ID
	UserId    int `json:"UserId" comment:"用户id"`          //用户ID
	TagId     int `json:"TagId"  comment:"标签id"  `        //标签ID
	common.ControlBy
}

func (s *UserTagInsertReq) GenerateUserPatent(c *models.UserTag) {
	s.UserId = c.UserId
	s.TagId = c.TagId
}
func NewUserTagInsert(userId, tagId int) *UserTagInsertReq {
	return &UserTagInsertReq{
		UserId: userId,
		TagId:  tagId,
	}
}
