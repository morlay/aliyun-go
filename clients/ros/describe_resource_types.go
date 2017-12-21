package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceTypesRequest struct {
	SupportStatus string `position:"Query" name:"SupportStatus"`
}

func (r DescribeResourceTypesRequest) Invoke(client *sdk.Client) (response *DescribeResourceTypesResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeResourceTypesRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeResourceTypes", "/resource_types", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeResourceTypesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeResourceTypesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeResourceTypesResponse struct {
}
