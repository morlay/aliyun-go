package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterScaledNodeRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *DescribeClusterScaledNodeRequest) Invoke(client *sdk.Client) (resp *DescribeClusterScaledNodeResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterScaledNode", "/clusters/[ClusterId]/scaled_nodes/", "", "")
	req.Method = "GET"

	resp = &DescribeClusterScaledNodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterScaledNodeResponse struct {
	responses.BaseResponse
}
