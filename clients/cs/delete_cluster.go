package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteClusterRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *DeleteClusterRequest) Invoke(client *sdk.Client) (resp *DeleteClusterResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DeleteCluster", "/clusters/[ClusterId]", "", "")
	req.Method = "DELETE"

	resp = &DeleteClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteClusterResponse struct {
	responses.BaseResponse
}
