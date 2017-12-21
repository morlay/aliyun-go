package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseClusterPublicConnectionRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId             string `position:"Query" name:"DBClusterId"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
}

func (r ReleaseClusterPublicConnectionRequest) Invoke(client *sdk.Client) (response *ReleaseClusterPublicConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReleaseClusterPublicConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "ReleaseClusterPublicConnection", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		ReleaseClusterPublicConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReleaseClusterPublicConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReleaseClusterPublicConnectionResponse struct {
	RequestId string
}
