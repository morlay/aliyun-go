package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiTrafficControlsRequest struct {
	StageName  string `position:"Query" name:"StageName"`
	GroupId    string `position:"Query" name:"GroupId"`
	ApiIds     string `position:"Query" name:"ApiIds"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
}

func (r DescribeApiTrafficControlsRequest) Invoke(client *sdk.Client) (response *DescribeApiTrafficControlsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeApiTrafficControlsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiTrafficControls", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeApiTrafficControlsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeApiTrafficControlsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeApiTrafficControlsResponse struct {
	RequestId          string
	TotalCount         int
	PageSize           int
	PageNumber         int
	ApiTrafficControls DescribeApiTrafficControlsApiTrafficControlItemList
}

type DescribeApiTrafficControlsApiTrafficControlItem struct {
	ApiId              string
	ApiName            string
	TrafficControlId   string
	TrafficControlName string
	BoundTime          string
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
