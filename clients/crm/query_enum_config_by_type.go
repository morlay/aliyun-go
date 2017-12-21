package crm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryEnumConfigByTypeRequest struct {
	Type string `position:"Query" name:"Type"`
}

func (r QueryEnumConfigByTypeRequest) Invoke(client *sdk.Client) (response *QueryEnumConfigByTypeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryEnumConfigByTypeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Crm", "2015-03-24", "QueryEnumConfigByType", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryEnumConfigByTypeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryEnumConfigByTypeResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryEnumConfigByTypeResponse struct {
	Success       bool
	ResultCode    string
	ResultMessage string
	Data          QueryEnumConfigByTypeBizCategoryList
}

type QueryEnumConfigByTypeBizCategory struct {
	EnumName  string
	EnumValue string
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
