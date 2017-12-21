package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateClusterRequest struct {
	requests.RoaRequest
}

func (req *CreateClusterRequest) Invoke(client *sdk.Client) (resp *CreateClusterResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "CreateCluster", "/clusters", "", "")
	req.Method = "POST"

	resp = &CreateClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateClusterResponse struct {
	responses.BaseResponse
}
