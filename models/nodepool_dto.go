package models

type NodePoolDto struct {
	Id   int32  `gorm:"column:id;type:int;primaryKey;<-:false"`
	Name string `gorm:"column:name;type:varchar(253);<-:false"`
}

func (NodePoolDto) TableName() string {
	return "cost.node_pool"
}
