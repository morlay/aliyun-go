package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSystemParametersRequest struct {
	requests.RpcRequest
}

func (req *DescribeSystemParametersRequest) Invoke(client *sdk.Client) (resp *DescribeSystemParametersResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeSystemParameters", "apigateway", "")
	resp = &DescribeSystemParametersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSystemParametersResponse struct {
	responses.BaseResponse
	RequestId    common.String
	SystemParams DescribeSystemParametersSystemParamItemList
}

type DescribeSystemParametersSystemParamItem struct {
	ParamName   common.String
	ParamType   common.String
	DemoValue   common.String
	Description common.String
}

type DescribeSystemParametersSystemParamItemList []DescribeSystemParametersSystemParamItem

func (list *DescribeSystemParametersSystemParamItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSystemParametersSystemParamItem)
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
