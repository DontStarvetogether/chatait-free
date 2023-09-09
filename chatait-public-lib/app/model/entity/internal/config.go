// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// Config is the golang structure for table c_config.
type Config struct {
	ConfigName string `orm:"config_name,primary" json:"configName"` // 配置参数名
	Title      string `orm:"title"               json:"title"`      // 标题
	Unit       string `orm:"unit"                json:"unit"`       // 单位
	InputType  int    `orm:"input_type"          json:"inputType"`  // 表单类型
	Options    string `orm:"options"             json:"options"`    // 参数配置的选项
	Value      string `orm:"value"               json:"value"`      // 配置值
	Type       string `orm:"type"                json:"type"`       // 类型
	Sort       int    `orm:"sort"                json:"sort"`       // 排序
	CreatedAt  int    `orm:"created_at"          json:"createdAt"`  // 创建时间
	UpdatedAt  int    `orm:"updated_at"          json:"updatedAt"`  // 更新时间
}