package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ScaleInClusterRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *ScaleInClusterRequest) Invoke(client *sdk.Client) (resp *ScaleInClusterResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "ScaleInCluster", "/clusters/[ClusterId]/scalein", "", "")
	req.Method = "POST"

	resp = &ScaleInClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type ScaleInClusterResponse struct {
	responses.BaseResponse
}
