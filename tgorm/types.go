package tgorm

type GormCommonField struct {
	ID         int64 `gorm:"primary_key;Column:id" json:"id"`
	IsDelete   int8  `gorm:"Column:is_delete" json:"is_delete"`
	CreateTime int64 `gorm:"Column:create_time" json:"create_time"`
	UpdateTime int64 `gorm:"Column:update_time" json:"update_time"`
	DeleteTime int64 `gorm:"Column:delete_time" json:"delete_time"`
}

type Dict struct {
	GormCommonField
	Type       string `gorm:"Column:type" json:"type"`
	ParentCode string `gorm:"Column:parent_code" json:"parent_code"`
	Code       string `gorm:"Column:code" json:"code"`
	Value      string `gorm:"Column:value" json:"value"`
	Remark     string `gorm:"Column:remark" json:"remark"`
}

func (Dict) TableName() string {
	return "t_dict"
}
