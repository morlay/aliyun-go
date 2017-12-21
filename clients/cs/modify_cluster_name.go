package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyClusterNameRequest struct {
	requests.RoaRequest
}

func (req *ModifyClusterNameRequest) Invoke(client *sdk.Client) (resp *ModifyClusterNameResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "ModifyClusterName", "/clusters/[ClusterId]/name/[ClusterName]", "", "")
	req.Method = "POST"

	resp = &ModifyClusterNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyClusterNameResponse struct {
	responses.BaseResponse
}
