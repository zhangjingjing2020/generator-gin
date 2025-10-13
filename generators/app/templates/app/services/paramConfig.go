package services

import (
	"<%= moduleName %>/app/models"
	"<%= moduleName %>/global"
	"errors"
)

type paramConfigService struct {
}

var ParamConfigService = new(paramConfigService)

// 查询配置信息
func (paramConfigService *paramConfigService) GetParamByKey(k string, subcompany_id string, mixer_no string) (pc models.ParamConfig, err error) {
	if global.App.DB == nil {
		return pc, errors.New("数据库连接错误")
	}
	err = global.App.DB.Where("status = ? AND param_key = ? AND subcompany_id = ? AND mixer_no = ?", 1, k, subcompany_id, mixer_no).First(&pc).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = errors.New("数据不存在")
		}
	}
	return
}

func (paramConfigService *paramConfigService) SetParamByKey(k, val string, subcompany_id string, mixer_no string) (pc models.ParamConfig, err error) {
	if global.App.DB == nil {
		return pc, errors.New("数据库连接错误")
	}
	res := global.App.DB.Where("status = ? AND param_key = ? AND subcompany_id = ? AND mixer_no = ?", 1, k, subcompany_id, mixer_no).First(&pc)

	if res.RowsAffected > 0 {
		pc.ParamValue = val
		global.App.DB.Save(&pc)
	}

	return
}

// 新增配置参数
func (paramConfigService *paramConfigService) AddParamByKey(k string, v string, subcompany_id string, mixer_no string) (pc models.ParamConfig, err error) {
	if global.App.DB == nil {
		return pc, errors.New("数据库连接错误")
	}
	pc = models.ParamConfig{ParamKey: k, ParamValue: v, SubcompanyId: subcompany_id, MixerNo: mixer_no, Status: 1}

	err = global.App.DB.Create(&pc).Error
	return
}
