package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRegionsRequest struct {
}

func (r DescribeRegionsRequest) Invoke(client *sdk.Client) (response *DescribeRegionsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeRegionsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DescribeRegions", "/regions", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeRegionsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRegionsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRegionsResponse struct {
}
