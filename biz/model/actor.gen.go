// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameActor = "actor"

// Actor mapped from table <actor>
type Actor struct {
	ID         int32          `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:作者id" json:"id,string"`
	Name       string         `gorm:"column:name;type:varchar(255);not null;comment:作者名" json:"name"`
	UpdateTime time.Time      `gorm:"column:update_time;type:datetime;comment:更新时间" json:"update_time"`
	CreateTime time.Time      `gorm:"column:create_time;type:datetime;comment:创建时间" json:"create_time"`
	IsDeleted  gorm.DeletedAt `gorm:"column:is_deleted;type:datetime;comment:删除时间" json:"is_deleted"`
}

// TableName Actor's table name
func (*Actor) TableName() string {
	return TableNameActor
}