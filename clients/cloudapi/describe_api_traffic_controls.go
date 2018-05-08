package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiTrafficControlsRequest struct {
	requests.RpcRequest
	StageName  string `position:"Query" name:"StageName"`
	GroupId    string `position:"Query" name:"GroupId"`
	ApiIds     string `position:"Query" name:"ApiIds"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
}

func (req *DescribeApiTrafficControlsRequest) Invoke(client *sdk.Client) (resp *DescribeApiTrafficControlsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiTrafficControls", "apigateway", "")
	resp = &DescribeApiTrafficControlsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiTrafficControlsResponse struct {
	responses.BaseResponse
	RequestId          common.String
	TotalCount         common.Integer
	PageSize           common.Integer
	PageNumber         common.Integer
	ApiTrafficControls DescribeApiTrafficControlsApiTrafficControlItemList
}

type DescribeApiTrafficControlsApiTrafficControlItem struct {
	ApiId              common.String
	ApiName            common.String
	TrafficControlId   common.String
	TrafficControlName common.String
	BoundTime          common.String
}

type DescribeApiTrafficControlsApiTrafficControlItemList []DescribeApiTrafficControlsApiTrafficControlItem

func (list *DescribeApiTrafficControlsApiTrafficControlItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiTrafficControlsApiTrafficControlItem)
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
