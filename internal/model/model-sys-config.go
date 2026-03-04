package model

import (
	"gorm.io/datatypes"
	"helay.net/go/utils/v3/dataType"
)

type SysConfig struct {
	Prop       string               `json:"prop" gorm:"primaryKey;type:varchar(128)"`
	Label      string               `json:"label" gorm:"type:varchar(128);not null"`
	Value      string               `json:"value" gorm:"type:text;not null;default:''"`
	Name       string               `json:"name" gorm:"type:varchar(128);not null;default:'el-input'"`
	Type       string               `json:"type" gorm:"type:varchar(128);not null;default:'text'"`
	Component  datatypes.JSONMap    `json:"component" gorm:"type:text"`
	CreateTime *dataType.CustomTime `json:"create_time,omitempty" gorm:"autoCreateTime:true;index;not null;default:current_timestamp;comment:记录创建时间" form:"-"`
	UpdateTime *dataType.CustomTime `json:"update_time,omitempty" gorm:"autoUpdateTime:true;comment:记录更新时间" form:"-"`
}

func (c *SysConfig) TableName() string {
	return "sys_config"
}
