// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProductCategory = "ProductCategory"

// ProductCategory mapped from table <ProductCategory>
type ProductCategory struct {
	CategoryID   int32  `gorm:"column:CategoryID;primaryKey;autoIncrement:true" json:"CategoryID"`
	CategoryName string `gorm:"column:CategoryName;not null" json:"CategoryName"`
	Slug         string `gorm:"column:Slug;not null" json:"Slug"`
}

// TableName ProductCategory's table name
func (*ProductCategory) TableName() string {
	return TableNameProductCategory
}
