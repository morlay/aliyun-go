package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourcesRequest struct {
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (r DescribeResourcesRequest) Invoke(client *sdk.Client) (response *DescribeResourcesResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeResourcesRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeResources", "/stacks/[StackName]/[StackId]/resources", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeResourcesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeResourcesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeResourcesResponse struct {
}
