package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterTokensRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *DescribeClusterTokensRequest) Invoke(client *sdk.Client) (resp *DescribeClusterTokensResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterTokens", "/clusters/[ClusterId]/tokens", "", "")
	req.Method = "GET"

	resp = &DescribeClusterTokensResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterTokensResponse struct {
	responses.BaseResponse
}
