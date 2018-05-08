package crm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryEnumConfigByTypeRequest struct {
	requests.RpcRequest
	Type string `position:"Query" name:"Type"`
}

func (req *QueryEnumConfigByTypeRequest) Invoke(client *sdk.Client) (resp *QueryEnumConfigByTypeResponse, err error) {
	req.InitWithApiInfo("Crm", "2015-03-24", "QueryEnumConfigByType", "", "")
	resp = &QueryEnumConfigByTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryEnumConfigByTypeResponse struct {
	responses.BaseResponse
	Success       bool
	ResultCode    common.String
	ResultMessage common.String
	Data          QueryEnumConfigByTypeBizCategoryList
}

type QueryEnumConfigByTypeBizCategory struct {
	EnumName  common.String
	EnumValue common.String
}

type QueryEnumConfigByTypeBizCategoryList []QueryEnumConfigByTypeBizCategory

func (list *QueryEnumConfigByTypeBizCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEnumConfigByTypeBizCategory)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
