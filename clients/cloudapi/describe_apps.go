package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAppsRequest struct {
	AppId      int64  `position:"Query" name:"AppId"`
	AppOwner   string `position:"Query" name:"AppOwner"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
}

func (r DescribeAppsRequest) Invoke(client *sdk.Client) (response *DescribeAppsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAppsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApps", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAppsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAppsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAppsResponse struct {
	RequestId  string
	TotalCount int
	PageSize   int
	PageNumber int
	Apps       DescribeAppsAppItemList
}

type DescribeAppsAppItem struct {
	AppId       int64
	AppName     string
	Description string
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
