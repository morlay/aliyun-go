package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiIpControlsRequest struct {
	requests.RpcRequest
	StageName  string `position:"Query" name:"StageName"`
	GroupId    string `position:"Query" name:"GroupId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	ApiIds     string `position:"Query" name:"ApiIds"`
}

func (req *DescribeApiIpControlsRequest) Invoke(client *sdk.Client) (resp *DescribeApiIpControlsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiIpControls", "apigateway", "")
	resp = &DescribeApiIpControlsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiIpControlsResponse struct {
	responses.BaseResponse
	RequestId     common.String
	TotalCount    common.Integer
	PageSize      common.Integer
	PageNumber    common.Integer
	ApiIpControls DescribeApiIpControlsApiIpControlItemList
}

type DescribeApiIpControlsApiIpControlItem struct {
	ApiId         common.String
	ApiName       common.String
	IpControlId   common.String
	IpControlName common.String
	BoundTime     common.String
}

type DescribeApiIpControlsApiIpControlItemList []DescribeApiIpControlsApiIpControlItem

func (list *DescribeApiIpControlsApiIpControlItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiIpControlsApiIpControlItem)
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
