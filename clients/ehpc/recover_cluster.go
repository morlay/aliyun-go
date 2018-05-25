package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RecoverClusterRequest struct {
	requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

func (req *RecoverClusterRequest) Invoke(client *sdk.Client) (resp *RecoverClusterResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "RecoverCluster", "ehs", "")
	resp = &RecoverClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type RecoverClusterResponse struct {
	responses.BaseResponse
	RequestId common.String
}
