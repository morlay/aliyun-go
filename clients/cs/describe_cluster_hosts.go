package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterHostsRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *DescribeClusterHostsRequest) Invoke(client *sdk.Client) (resp *DescribeClusterHostsResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterHosts", "/clusters/[ClusterId]/hosts", "", "")
	req.Method = "GET"

	resp = &DescribeClusterHostsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterHostsResponse struct {
	responses.BaseResponse
}
