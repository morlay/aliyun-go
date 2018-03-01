package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAgilityTunnelCertsRequest struct {
	requests.RoaRequest
	Token string `position:"Path" name:"Token"`
}

func (req *DescribeAgilityTunnelCertsRequest) Invoke(client *sdk.Client) (resp *DescribeAgilityTunnelCertsResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeAgilityTunnelCerts", "/agility/[Token]/agent_certs", "", "")
	req.Method = "GET"

	resp = &DescribeAgilityTunnelCertsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAgilityTunnelCertsResponse struct {
	responses.BaseResponse
}
