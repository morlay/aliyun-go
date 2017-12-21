package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceDetailRequest struct {
	StackName    string `position:"Path" name:"StackName"`
	StackId      string `position:"Path" name:"StackId"`
	ResourceName string `position:"Path" name:"ResourceName"`
}

func (r DescribeResourceDetailRequest) Invoke(client *sdk.Client) (response *DescribeResourceDetailResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeResourceDetailRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeResourceDetail", "/stacks/[StackName]/[StackId]/resources/[ResourceName]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeResourceDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeResourceDetailResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeResourceDetailResponse struct {
}
