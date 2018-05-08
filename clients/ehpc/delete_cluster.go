package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteClusterRequest struct {
	requests.RpcRequest
	ReleaseInstance string `position:"Query" name:"ReleaseInstance"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DeleteClusterRequest) Invoke(client *sdk.Client) (resp *DeleteClusterResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "DeleteCluster", "ehs", "")
	resp = &DeleteClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteClusterResponse struct {
	responses.BaseResponse
	RequestId common.String
}
