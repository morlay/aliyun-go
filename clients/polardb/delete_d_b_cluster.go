package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDBClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteDBClusterRequest) Invoke(client *sdk.Client) (resp *DeleteDBClusterResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DeleteDBCluster", "polardb", "")
	resp = &DeleteDBClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDBClusterResponse struct {
	responses.BaseResponse
	RequestId string
}
