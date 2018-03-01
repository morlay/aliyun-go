package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddAgilityClusterRequest struct {
	requests.RoaRequest
}

func (req *AddAgilityClusterRequest) Invoke(client *sdk.Client) (resp *AddAgilityClusterResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "AddAgilityCluster", "/add_agility_cluster", "", "")
	req.Method = "POST"

	resp = &AddAgilityClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddAgilityClusterResponse struct {
	responses.BaseResponse
}
