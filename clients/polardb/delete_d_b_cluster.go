package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDBClusterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteDBClusterRequest) Invoke(client *sdk.Client) (response *DeleteDBClusterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDBClusterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DeleteDBCluster", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDBClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDBClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDBClusterResponse struct {
	RequestId string
}
