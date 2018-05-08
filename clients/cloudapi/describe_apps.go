package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAppsRequest struct {
	requests.RpcRequest
	AppId      int64  `position:"Query" name:"AppId"`
	AppOwner   string `position:"Query" name:"AppOwner"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
}

func (req *DescribeAppsRequest) Invoke(client *sdk.Client) (resp *DescribeAppsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApps", "apigateway", "")
	resp = &DescribeAppsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAppsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageSize   common.Integer
	PageNumber common.Integer
	Apps       DescribeAppsAppItemList
}

type DescribeAppsAppItem struct {
	AppId       common.Long
	AppName     common.String
	Description common.String
}

type DescribeAppsAppItemList []DescribeAppsAppItem

func (list *DescribeAppsAppItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAppsAppItem)
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
