package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ScaleClusterRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *ScaleClusterRequest) Invoke(client *sdk.Client) (resp *ScaleClusterResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "ScaleCluster", "/clusters/[ClusterId]", "", "")
	req.Method = "PUT"

	resp = &ScaleClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type ScaleClusterResponse struct {
	responses.BaseResponse
}
