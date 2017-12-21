package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceTypesRequest struct {
	requests.RoaRequest
	SupportStatus string `position:"Query" name:"SupportStatus"`
}

func (req *DescribeResourceTypesRequest) Invoke(client *sdk.Client) (resp *DescribeResourceTypesResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeResourceTypes", "/resource_types", "", "")
	req.Method = "GET"

	resp = &DescribeResourceTypesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeResourceTypesResponse struct {
	responses.BaseResponse
}
