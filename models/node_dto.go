package models

type NodeDto struct {
	Id         int32        `gorm:"column:id;type:int;primaryKey;<-:false"`
	Name       string       `gorm:"column:name;type:varchar(253);<-:false"`
	NodePoolId *int32       `gorm:"column:pool_id;type:int;<-:false"`
	NodePool   *NodePoolDto `gorm:"foreignKey:NodePoolId;<-:false"`
}

func (NodeDto) TableName() string {
	return "cost.nodes"
}
