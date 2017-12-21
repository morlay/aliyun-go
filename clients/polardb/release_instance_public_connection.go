package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseInstancePublicConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
}

func (req *ReleaseInstancePublicConnectionRequest) Invoke(client *sdk.Client) (resp *ReleaseInstancePublicConnectionResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ReleaseInstancePublicConnection", "polardb", "")
	resp = &ReleaseInstancePublicConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseInstancePublicConnectionResponse struct {
	responses.BaseResponse
	RequestId string
}
