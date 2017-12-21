package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FailoverDBClusterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	TargetDBInstanceId   string `position:"Query" name:"TargetDBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r FailoverDBClusterRequest) Invoke(client *sdk.Client) (response *FailoverDBClusterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		FailoverDBClusterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "FailoverDBCluster", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		FailoverDBClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.FailoverDBClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type FailoverDBClusterResponse struct {
	RequestId string
}
