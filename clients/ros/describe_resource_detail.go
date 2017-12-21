package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceDetailRequest struct {
	requests.RoaRequest
	StackName    string `position:"Path" name:"StackName"`
	StackId      string `position:"Path" name:"StackId"`
	ResourceName string `position:"Path" name:"ResourceName"`
}

func (req *DescribeResourceDetailRequest) Invoke(client *sdk.Client) (resp *DescribeResourceDetailResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeResourceDetail", "/stacks/[StackName]/[StackId]/resources/[ResourceName]", "", "")
	req.Method = "GET"

	resp = &DescribeResourceDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeResourceDetailResponse struct {
	responses.BaseResponse
}
