// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameNotification = "Notification"

// Notification mapped from table <Notification>
type Notification struct {
	NotificationID int32     `gorm:"column:NotificationID;primaryKey;autoIncrement:true" json:"NotificationID"`
	UserID         int32     `gorm:"column:UserID;not null" json:"UserID"`
	Content        string    `gorm:"column:Content;not null" json:"Content"`
	Link           string    `gorm:"column:Link;not null" json:"Link"`
	CreatedAt      time.Time `gorm:"column:CreatedAt" json:"CreatedAt"`
}

// TableName Notification's table name
func (*Notification) TableName() string {
	return TableNameNotification
}
