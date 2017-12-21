package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeStacksRequest struct {
	requests.RoaRequest
	Status     string `position:"Query" name:"Status"`
	Name       string `position:"Query" name:"Name"`
	StackId    string `position:"Query" name:"StackId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeStacksRequest) Invoke(client *sdk.Client) (resp *DescribeStacksResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeStacks", "/stacks", "", "")
	req.Method = "GET"

	resp = &DescribeStacksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeStacksResponse struct {
	responses.BaseResponse
}
