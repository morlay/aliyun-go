package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RestartClusterRequest struct {
	requests.RpcRequest
	RollingRestart        string `position:"Query" name:"RollingRestart"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	UpgradedHostGroupOnly string `position:"Query" name:"UpgradedHostGroupOnly"`
	ClusterId             string `position:"Query" name:"ClusterId"`
}

func (req *RestartClusterRequest) Invoke(client *sdk.Client) (resp *RestartClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RestartCluster", "", "")
	resp = &RestartClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type RestartClusterResponse struct {
	responses.BaseResponse
	RequestId common.String
	ClusterId common.String
}
