package labdb

type GormCommonField struct {
	ID         int64 `gorm:"primary_key;Column:id" json:"id"`
	IsDelete   int8  `gorm:"Column:is_delete" json:"is_delete"`
	CreateTime int64 `gorm:"Column:create_time" json:"create_time"`
	UpdateTime int64 `gorm:"Column:update_time" json:"update_time"`
	DeleteTime int64 `gorm:"Column:delete_time" json:"delete_time"`
}

//TaskConfig 健康任务奖励配置
type TaskConfig struct {
	GormCommonField
	GroupID     int    `json:"group_id"`
	GroupTitle  string `json:"group_title"`
	GroupStatus int    `json:"group_status"`
	TaskID      int    `json:"task_id"`
	TaskKind    int    `json:"task_kind"`
	TaskTitle   string `json:"task_title"`
	RewardRule  int    `json:"reward_rule"`
	RewardType  int    `json:"reward_type"`
	RewardValue string `json:"reward_value"`
}

//TableName 表名
func (TaskConfig) TableName() string {
	return "t_task_config"
}
