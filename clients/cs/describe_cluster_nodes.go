package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterNodesRequest struct {
	requests.RoaRequest
	PageSize   string `position:"Query" name:"PageSize"`
	ClusterId  string `position:"Path" name:"ClusterId"`
	PageNumber string `position:"Query" name:"PageNumber"`
}

func (req *DescribeClusterNodesRequest) Invoke(client *sdk.Client) (resp *DescribeClusterNodesResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterNodes", "/clusters/[ClusterId]/nodes", "", "")
	req.Method = "GET"

	resp = &DescribeClusterNodesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterNodesResponse struct {
	responses.BaseResponse
}
