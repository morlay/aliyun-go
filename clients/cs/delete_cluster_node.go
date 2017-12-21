package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteClusterNodeRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
	Ip        string `position:"Path" name:"Ip"`
	Force     string `position:"Query" name:"Force"`
}

func (req *DeleteClusterNodeRequest) Invoke(client *sdk.Client) (resp *DeleteClusterNodeResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DeleteClusterNode", "/clusters/[ClusterId]/ip/[Ip]", "", "")
	req.Method = "DELETE"

	resp = &DeleteClusterNodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteClusterNodeResponse struct {
	responses.BaseResponse
}
