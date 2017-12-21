package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTemplateRequest struct {
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (r DescribeTemplateRequest) Invoke(client *sdk.Client) (response *DescribeTemplateResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeTemplateRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeTemplate", "/stacks/[StackName]/[StackId]/template", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTemplateResponse struct {
}
