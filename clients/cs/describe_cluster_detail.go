package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterDetailRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *DescribeClusterDetailRequest) Invoke(client *sdk.Client) (resp *DescribeClusterDetailResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterDetail", "/clusters/[ClusterId]", "", "")
	req.Method = "GET"

	resp = &DescribeClusterDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterDetailResponse struct {
	responses.BaseResponse
}
