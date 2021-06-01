package models

import "time"

type ContainerDto struct {
	ContainerId            string    `gorm:"column:container_id;type:varchar(253);primaryKey;<-:false"`
	ContainerName          string    `gorm:"column:container_name;type:varchar(253);<-:false"`
	PodName                string    `gorm:"column:pod_name;type:varchar(253);<-:false"`
	ApplicationName        string    `gorm:"column:application_name;type:varchar(253);<-:false"`
	EnvironmentName        string    `gorm:"column:environment_name;type:varchar(253);<-:false"`
	ComponentName          string    `gorm:"column:component_name;type:varchar(253);<-:false"`
	WBS                    string    `gorm:"column:wbs;type:varchar(253);<-:false"`
	StartedAt              time.Time `gorm:"column:started_at;type:datetimeoffset(0);<-:false"`
	LastKnownRunningAt     time.Time `gorm:"column:last_known_running_at;type:datetimeoffset(0);<-:false"`
	CpuRequestedMillicores int64     `gorm:"column:cpu_request_millicores;type:bigint;<-:false"`
	MemoryRequestedBytes   int64     `gorm:"column:memory_request_bytes;type:bigint;<-:false"`
	NodeId                 int32     `gorm:"column:node_id;type:int;<-:false"`
	Node                   *NodeDto  `gorm:"foreignKey:NodeId"`
}

func (ContainerDto) TableName() string {
	return "cost.containers"
}
