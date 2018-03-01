package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAgilityTunnelAgentInfoRequest struct {
	requests.RoaRequest
	Token string `position:"Path" name:"Token"`
}

func (req *DescribeAgilityTunnelAgentInfoRequest) Invoke(client *sdk.Client) (resp *DescribeAgilityTunnelAgentInfoResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeAgilityTunnelAgentInfo", "/agility/[Token]/agent_info", "", "")
	req.Method = "GET"

	resp = &DescribeAgilityTunnelAgentInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAgilityTunnelAgentInfoResponse struct {
	responses.BaseResponse
}
