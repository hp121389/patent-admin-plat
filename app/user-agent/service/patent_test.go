package service

import (
	"go-admin/app/user-agent/models"
	common "go-admin/common/models"
	"testing"
)

func TestFindInventorsAndRelationsFromPatents(t *testing.T) {
	pts := make([]models.Patent, 0)
	pts = append(pts, models.Patent{
		PatentId:         1,
		PNM:              "1",
		PatentProperties: "{\"patentId\":11,\"TI\":\"基于11的专利\",\"PNM\":\"111\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of T0Software\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author011\",\"CLS\":\"有权\",\"CreateBy\":1,\"UpdateBy\":0, \"INN\":\"沈家琦;黄涛;胡泊;sjq;ht;ysd\"}",
		ControlBy:        common.ControlBy{},
		ModelTime:        common.ModelTime{},
	})

	pts = append(pts, models.Patent{
		PatentId:         2,
		PNM:              "2",
		PatentProperties: "{\"patentId\":11,\"TI\":\"基于11的专利\",\"PNM\":\"111\",\"AD\":\"2022-10-18 18:49:53\",\"PD\":\"2022-10-20 18:49:53\",\"CL\":\"a patent of T0Software\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author011\",\"CLS\":\"有权\",\"CreateBy\":1,\"UpdateBy\":0, \"INN\":\"沈家琦;黄涛;sjq;ht\"}",
		ControlBy:        common.ControlBy{},
		ModelTime:        common.ModelTime{},
	})

	//i, r, err := Fin(pts, 200)
	//assert.NoError(t, err)
	//fmt.Println(i)
	//fmt.Println(r)
}
