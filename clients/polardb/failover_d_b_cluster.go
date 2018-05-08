package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FailoverDBClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	TargetDBInstanceId   string `position:"Query" name:"TargetDBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *FailoverDBClusterRequest) Invoke(client *sdk.Client) (resp *FailoverDBClusterResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "FailoverDBCluster", "polardb", "")
	resp = &FailoverDBClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type FailoverDBClusterResponse struct {
	responses.BaseResponse
	RequestId common.String
}
