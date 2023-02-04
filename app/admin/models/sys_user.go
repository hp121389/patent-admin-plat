package models

import (
	"go-admin/common/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string `json:"username" gorm:"size:64;comment:用户名"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`
	RoleId   int    `json:"roleId" gorm:"size:20;comment:角色ID"`
	Salt     string `json:"-" gorm:"size:255;comment:加盐"`
	Avatar   string `json:"avatar" gorm:"size:255;comment:头像"`
	Sex      string `json:"sex" gorm:"size:255;comment:性别"`
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`
	Remark   string `json:"remark" gorm:"size:255;comment:备注"`
	Bio      string `json:"bio" gorm:"size:1024;comment:简介"`
	Interest string `json:"interest" gorm:"size:1024;comment:研究兴趣"`
	DepartID int    `json:"departID" gorm:"size:128;comment:单位"`
	Status   string `json:"status" gorm:"size:4;comment:状态"`
	RoleIds  []int  `json:"roleIds" gorm:"-"` //忽略该字段，- 表示无读写
	models.ControlBy
	models.ModelTime
	//嵌入结构体：先写好models然后嵌入，等效于嵌入的models也在初始化数据表时生效

	Dept SysDept `json:"dept"  gorm:"-"`
}

func (e *SysUser) TableName() string {
	return "sys_user"
}

func (e *SysUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysUser) GetId() interface{} {
	return e.UserId
}

/*
GORM 允许用户定义的钩子有
创建记录时将调用这些钩子方法 BeforeSave,BeforeCreate,/更新db/,AfterSave, AfterCreate 提交或回滚事务
更新记录时将调用这些钩子方法 BeforeSave,BeforeUpdate,/更新db/,AfterUpdate,AfterSave 提交或回滚事务
删除记录时将调用这些钩子方法 BeforeDelete,/更新db/,AfterDelete 提交或回滚事务
查询记录时将调用这些钩子方法 /从 db 中加载数据/,Preloading (eager loading),AfterFind

如果您已经为模型定义了指定的方法，它会在创建、更新、查询、删除时自动被调用。如果任何回调返回错误，GORM 将停止后续的操作并回滚事务。
*/
//钩子方法的函数签名应该是 func(*gorm.DB) error

//加密

func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *SysUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}

func (e *SysUser) AfterFind(_ *gorm.DB) error {
	e.RoleIds = []int{e.RoleId}
	return nil
}
