package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchQueryScaleFlavorsResponse struct {
	// 计算规格列表对象。

	ComputeFlavorGroups *[]Computes `json:"compute_flavor_groups,omitempty"`
	HttpStatusCode      int         `json:"-"`
}

func (o SearchQueryScaleFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQueryScaleFlavorsResponse struct{}"
	}

	return strings.Join([]string{"SearchQueryScaleFlavorsResponse", string(data)}, " ")
}
