package models

// 参数配置表
type ParamConfig struct {
	ID
	ParamKey     string `gorm:"column:param_key" json:"param_key"`
	ParamValue   string `gorm:"column:param_value" json:"param_value"`
	SubcompanyId string `gorm:"column:subcompany_id" json:"subcompany_id"`
	MixerNo      string `gorm:"column:mixer_no" json:"mixer_no"`
	Status       int8   `gorm:"column:status" json:"status"`
	Timestamps
	SoftDeletes
}

func (ParamConfig) TableName() string {
	return "param_config"
}
