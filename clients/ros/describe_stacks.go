package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeStacksRequest struct {
	Status     string `position:"Query" name:"Status"`
	Name       string `position:"Query" name:"Name"`
	StackId    string `position:"Query" name:"StackId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (r DescribeStacksRequest) Invoke(client *sdk.Client) (response *DescribeStacksResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeStacksRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeStacks", "/stacks", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeStacksResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeStacksResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeStacksResponse struct {
}
