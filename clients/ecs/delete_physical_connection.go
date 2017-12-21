package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePhysicalConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeletePhysicalConnectionRequest) Invoke(client *sdk.Client) (response *DeletePhysicalConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeletePhysicalConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeletePhysicalConnection", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeletePhysicalConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeletePhysicalConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeletePhysicalConnectionResponse struct {
	RequestId string
}
