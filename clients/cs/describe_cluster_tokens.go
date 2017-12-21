package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterTokensRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (r DescribeClusterTokensRequest) Invoke(client *sdk.Client) (response *DescribeClusterTokensResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeClusterTokensRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterTokens", "/clusters/[ClusterId]/tokens", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeClusterTokensResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeClusterTokensResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeClusterTokensResponse struct {
}
