package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StopClusterRequest struct {
	requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

func (req *StopClusterRequest) Invoke(client *sdk.Client) (resp *StopClusterResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "StopCluster", "ehs", "")
	resp = &StopClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopClusterResponse struct {
	responses.BaseResponse
	RequestId common.String
}
