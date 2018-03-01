package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetClusterNodeRequest struct {
	requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	ClusterId  string `position:"Path" name:"ClusterId"`
}

func (req *ResetClusterNodeRequest) Invoke(client *sdk.Client) (resp *ResetClusterNodeResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "ResetClusterNode", "/clusters/[ClusterId]/instances/[InstanceId]/reset", "", "")
	req.Method = "POST"

	resp = &ResetClusterNodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetClusterNodeResponse struct {
	responses.BaseResponse
}
