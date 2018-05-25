package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StartClusterRequest struct {
	requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

func (req *StartClusterRequest) Invoke(client *sdk.Client) (resp *StartClusterResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "StartCluster", "ehs", "")
	resp = &StartClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartClusterResponse struct {
	responses.BaseResponse
	RequestId common.String
}
