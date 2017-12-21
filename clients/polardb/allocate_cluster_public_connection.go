package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AllocateClusterPublicConnectionRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	DBClusterId            string `position:"Query" name:"DBClusterId"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (r AllocateClusterPublicConnectionRequest) Invoke(client *sdk.Client) (response *AllocateClusterPublicConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AllocateClusterPublicConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "AllocateClusterPublicConnection", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		AllocateClusterPublicConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AllocateClusterPublicConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type AllocateClusterPublicConnectionResponse struct {
	RequestId        string
	DBInstanceId     string
	ConnectionString string
	IPType           string
	Port             string
}
