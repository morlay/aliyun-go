package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterDetailRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (r DescribeClusterDetailRequest) Invoke(client *sdk.Client) (response *DescribeClusterDetailResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeClusterDetailRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterDetail", "/clusters/[ClusterId]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeClusterDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeClusterDetailResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeClusterDetailResponse struct {
}
