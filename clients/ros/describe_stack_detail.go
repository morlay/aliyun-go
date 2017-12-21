package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeStackDetailRequest struct {
	requests.RoaRequest
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (req *DescribeStackDetailRequest) Invoke(client *sdk.Client) (resp *DescribeStackDetailResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeStackDetail", "/stacks/[StackName]/[StackId]", "", "")
	req.Method = "GET"

	resp = &DescribeStackDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeStackDetailResponse struct {
	responses.BaseResponse
}
