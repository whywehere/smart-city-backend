// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEvent = "event"

// Event mapped from table <event>
type Event struct {
	ID       int32     `gorm:"column:id;primaryKey" json:"id"`
	State    int32     `gorm:"column:state" json:"state"`
	Time     time.Time `gorm:"column:time" json:"time"`
	Type     int32     `gorm:"column:type;comment:1-交通事故，2-犯罪事件， 3-火灾事故" json:"type"` // 1-交通事故，2-犯罪事件， 3-火灾事故
	Position string    `gorm:"column:position" json:"position"`
}

// TableName Event's table name
func (*Event) TableName() string {
	return TableNameEvent
}
