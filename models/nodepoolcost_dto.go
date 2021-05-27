package models

import "time"

type NodePoolCostDto struct {
	Id         int32       `gorm:"column:id;type:int;primaryKey;<-:false"`
	Cost       float64     `gorm:"column:cost;type:int;<-:false"`
	Currency   string      `gorm:"column:cost_currency;type:char(3);<-:false"`
	FromDate   time.Time   `gorm:"column:from_date;type:datetimeoffset(0);<-:false"`
	ToDate     time.Time   `gorm:"column:to_date;type:datetimeoffset(0);<-:false"`
	NodePoolId int32       `gorm:"column:pool_id;type:int;<-:false"`
	NodePool   NodePoolDto `gorm:"foreignKey:NodePoolId;<-:false"`
}

func (NodePoolCostDto) TableName() string {
	return "cost.node_pool_cost"
}
