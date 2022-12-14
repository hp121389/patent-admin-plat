package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/user-agent/models"
)

type Node struct {
	service.Service
}

// GetPatentIdByPackageId 通过PackageId获得PatentId
func (e *Node) GetNodes(list *[]models.Node, c int) error {
	var err error
	var data models.Node
	var data2 models.Node
	err = e.Orm.Model(&data).Where("node_id=?", c).
		Find(&data2).Limit(-1).Offset(-1).Error
	*list = append(*list, data2)
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	//fmt.Println(data2)
	return nil
}
