package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseClusterPublicConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId             string `position:"Query" name:"DBClusterId"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
}

func (req *ReleaseClusterPublicConnectionRequest) Invoke(client *sdk.Client) (resp *ReleaseClusterPublicConnectionResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ReleaseClusterPublicConnection", "polardb", "")
	resp = &ReleaseClusterPublicConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseClusterPublicConnectionResponse struct {
	responses.BaseResponse
	RequestId string
}
