package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourcesRequest struct {
	requests.RoaRequest
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (req *DescribeResourcesRequest) Invoke(client *sdk.Client) (resp *DescribeResourcesResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeResources", "/stacks/[StackName]/[StackId]/resources", "", "")
	req.Method = "GET"

	resp = &DescribeResourcesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeResourcesResponse struct {
	responses.BaseResponse
}
