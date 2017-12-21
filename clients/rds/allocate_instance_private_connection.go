package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AllocateInstancePrivateConnectionRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (r AllocateInstancePrivateConnectionRequest) Invoke(client *sdk.Client) (response *AllocateInstancePrivateConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AllocateInstancePrivateConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "AllocateInstancePrivateConnection", "rds", "")

	resp := struct {
		*responses.BaseResponse
		AllocateInstancePrivateConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AllocateInstancePrivateConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type AllocateInstancePrivateConnectionResponse struct {
	RequestId string
}
