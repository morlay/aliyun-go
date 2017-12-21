package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeStackDetailRequest struct {
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (r DescribeStackDetailRequest) Invoke(client *sdk.Client) (response *DescribeStackDetailResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeStackDetailRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeStackDetail", "/stacks/[StackName]/[StackId]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeStackDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeStackDetailResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeStackDetailResponse struct {
}
