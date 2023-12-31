// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameInvoiceDetail = "InvoiceDetail"

// InvoiceDetail mapped from table <InvoiceDetail>
type InvoiceDetail struct {
	ProductID     int32   `gorm:"column:ProductID;primaryKey" json:"ProductID"`
	InvoiceID     int32   `gorm:"column:InvoiceID;primaryKey" json:"InvoiceID"`
	Quantity      int32   `gorm:"column:Quantity;not null" json:"Quantity"`
	Price         int32   `gorm:"column:Price;not null" json:"Price"`
	Discount      float64 `gorm:"column:Discount;not null" json:"Discount"`
	AfterDiscount int32   `gorm:"column:AfterDiscount;not null" json:"AfterDiscount"`
}

// TableName InvoiceDetail's table name
func (*InvoiceDetail) TableName() string {
	return TableNameInvoiceDetail
}
