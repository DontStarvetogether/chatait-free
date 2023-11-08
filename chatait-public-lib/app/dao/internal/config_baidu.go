// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ConfigBaiduDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type ConfigBaiduDao struct {
	gmvc.M                     // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB             // DB is the raw underlying database management object.
	Table   string             // Table is the table name of the DAO.
	Columns configBaiduColumns // Columns contains all the columns of Table that for convenient usage.
}

// ConfigBaiduColumns defines and stores column names for table c_config_baidu.
type configBaiduColumns struct {
	Id        string // ID
	Title     string // 标题
	ApiKey    string //
	SecretKey string //
	Status    string // 启用状态
	CallNum   string // 调用次数
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

func NewConfigBaiduDao() *ConfigBaiduDao {
	return &ConfigBaiduDao{
		M:     g.DB("default").Model("c_config_baidu").Safe(),
		DB:    g.DB("default"),
		Table: "c_config_baidu",
		Columns: configBaiduColumns{
			Id:        "id",
			Title:     "title",
			ApiKey:    "api_key",
			SecretKey: "secret_key",
			Status:    "status",
			CallNum:   "call_num",
			CreatedAt: "created_at",
			UpdatedAt: "updated_at",
		},
	}
}
