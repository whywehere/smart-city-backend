// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSource = "source"

// Source mapped from table <source>
type Source struct {
	ID   int32   `gorm:"column:id;primaryKey" json:"id"`
	Type int32   `gorm:"column:type;primaryKey" json:"type"`
	Num  float32 `gorm:"column:num" json:"num"`
}

// TableName Source's table name
func (*Source) TableName() string {
	return TableNameSource
}