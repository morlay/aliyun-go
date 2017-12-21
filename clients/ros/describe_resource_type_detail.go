package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceTypeDetailRequest struct {
	TypeName string `position:"Path" name:"TypeName"`
}

func (r DescribeResourceTypeDetailRequest) Invoke(client *sdk.Client) (response *DescribeResourceTypeDetailResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeResourceTypeDetailRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeResourceTypeDetail", "/resource_types/[TypeName]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeResourceTypeDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeResourceTypeDetailResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeResourceTypeDetailResponse struct {
}
