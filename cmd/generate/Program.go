package main

import (
	"gorm.io/gorm"
	"time"
)

type Program struct {
	ID          int32          `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	Title       string         `gorm:"column:title;type:varchar(255)" json:"title"`
	View        string         `gorm:"column:view;type:varchar(255)" json:"view"`
	Actors      string         `gorm:"column:actors;type:varchar(255)" json:"actors"`
	ActorNum    int32          `gorm:"column:actor_num;type:int(11)" json:"actor_num"`
	CategorieID int32          `gorm:"column:categorie_id;type:int(11) unsigned" json:"categorie_id"`
	UpdeateTine time.Time      `gorm:"column:updeate_tine;type:datetime" json:"updeate_tine"`
	CreateTime  time.Time      `gorm:"column:create_time;type:datetime" json:"create_time"`
	IsDeleted   gorm.DeletedAt `gorm:"column:is_deleted;type:tinyint(4);not null" json:"is_deleted"`
}
