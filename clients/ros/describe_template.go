package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTemplateRequest struct {
	requests.RoaRequest
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (req *DescribeTemplateRequest) Invoke(client *sdk.Client) (resp *DescribeTemplateResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeTemplate", "/stacks/[StackName]/[StackId]/template", "", "")
	req.Method = "GET"

	resp = &DescribeTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTemplateResponse struct {
	responses.BaseResponse
}
