package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterNodeInfoRequest struct {
	requests.RoaRequest
	Token string `position:"Path" name:"Token"`
}

func (req *DescribeClusterNodeInfoRequest) Invoke(client *sdk.Client) (resp *DescribeClusterNodeInfoResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterNodeInfo", "/token/[Token]/node_info", "", "")
	req.Method = "GET"

	resp = &DescribeClusterNodeInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterNodeInfoResponse struct {
	responses.BaseResponse
}
