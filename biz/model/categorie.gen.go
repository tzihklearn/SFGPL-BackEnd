// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCategorie = "categorie"

// Categorie mapped from table <categorie>
type Categorie struct {
	ID         int32          `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:分类id" json:"id,string"`
	Name       string         `gorm:"column:name;type:varchar(255);comment:分类名" json:"name"`
	UpdateTime time.Time      `gorm:"column:update_time;type:datetime;comment:更新时间" json:"update_time"`
	CreateTime time.Time      `gorm:"column:create_time;type:datetime;comment:创建时间" json:"create_time"`
	IsDeleted  gorm.DeletedAt `gorm:"column:is_deleted;type:datetime;comment:删除时间" json:"is_deleted"`
}

// TableName Categorie's table name
func (*Categorie) TableName() string {
	return TableNameCategorie
}
