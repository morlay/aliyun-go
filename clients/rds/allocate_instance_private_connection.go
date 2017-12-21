package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AllocateInstancePrivateConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (req *AllocateInstancePrivateConnectionRequest) Invoke(client *sdk.Client) (resp *AllocateInstancePrivateConnectionResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "AllocateInstancePrivateConnection", "rds", "")
	resp = &AllocateInstancePrivateConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type AllocateInstancePrivateConnectionResponse struct {
	responses.BaseResponse
	RequestId string
}
