package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SetPostgresqlDbUserPwdRequest struct {
	// 语言

	XLanguage *SetPostgresqlDbUserPwdRequestXLanguage `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *DbUserPwdRequest `json:"body,omitempty"`
}

func (o SetPostgresqlDbUserPwdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPostgresqlDbUserPwdRequest struct{}"
	}

	return strings.Join([]string{"SetPostgresqlDbUserPwdRequest", string(data)}, " ")
}

type SetPostgresqlDbUserPwdRequestXLanguage struct {
	value string
}

type SetPostgresqlDbUserPwdRequestXLanguageEnum struct {
	ZH_CN SetPostgresqlDbUserPwdRequestXLanguage
	EN_US SetPostgresqlDbUserPwdRequestXLanguage
}

func GetSetPostgresqlDbUserPwdRequestXLanguageEnum() SetPostgresqlDbUserPwdRequestXLanguageEnum {
	return SetPostgresqlDbUserPwdRequestXLanguageEnum{
		ZH_CN: SetPostgresqlDbUserPwdRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SetPostgresqlDbUserPwdRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SetPostgresqlDbUserPwdRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetPostgresqlDbUserPwdRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
